package service

import (
	"brasa/db"
	"brasa/internal/receivable/storage"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateReceivable(c *gin.Context) {
	var newReceivable db.Receivable

	if err := c.ShouldBindJSON(&newReceivable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := storage.CreateReceivable(&newReceivable); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Receivable created successfully!", "receivable": newReceivable})
}

func GetReceivableByID(c *gin.Context) {
	id := c.Param("id")
	receivable, err := storage.GetReceivableByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receivable not found"})
		return
	}

	c.JSON(http.StatusOK, receivable)
}

func GetAllReceivables(c *gin.Context) {
	receivables, err := storage.GetAllReceivables()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receivables)
}

func DebtorPaymentAndDistribution(c *gin.Context, contractID string, paymentAmount float64) {
	contract, err := storage.GetContractByID(contractID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching contract"})
		return
	}

	if contract.Paid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment has already been made for this contract"})
		return
	}

	tokens, err := storage.GetAllTokens(contractID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tokens"})
		return
	}

	if len(tokens) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No tokens found for this contract"})
		return
	}

	totalTokens := float64(0)
	for _, token := range tokens {
		totalTokens += float64(token.Quantity)
	}

	if totalTokens == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No tokens issued"})
		return
	}

	for _, token := range tokens {
		proportion := float64(token.Quantity) / totalTokens
		amountToReceive := proportion * paymentAmount

		err := storage.RegisterDistribution(token.InvestorID, amountToReceive, token.ContractID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering distribution"})
			return
		}
	}

	contract.Paid = true
	contract.PaymentDate = time.Now()
	err = storage.UpdateContract(contract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating contract"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment successfully processed and distribution completed"})
}
