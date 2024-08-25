package service

import (
	"brasa/db"
	"brasa/internal/receivable/storage"
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{"message": "Recebível criado com sucesso!", "recebível": newReceivable})
}

func GetReceivableByID(c *gin.Context) {
	id := c.Param("id")
	receivable, err := storage.GetReceivableByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recebível não encontrado"})
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
