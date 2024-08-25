package service

import (
	"brasa/db"
	inter "brasa/internal/tokens/interface"
	"brasa/internal/tokens/storage"
	"brasa/package/lumx"
	lInterface "brasa/package/lumx/interface"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context, input inter.TokenCreateInput) {

	tokenQuantity := calculateTokenAmount(input.AssetValue)

	newInput := lInterface.TokenTypeRequest{
		Traits:      *input.Traits,
		MaxSupply:   tokenQuantity,
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
	}
	res, err := lumx.CreateTokenType(input.ContractId, newInput)
	if err != nil {
		c.Set("Error", "Erro na API Lumx")
		c.Status(http.StatusNotAcceptable)
		return
	}

	inputDB := inter.SaveToken{
		UriNumber:    uint64(res.UriNumber),
		ContractID:   res.ContractID,
		AssetID:      input.AssetID,
		NominalValue: input.NominalValue,
		ExpireDate:   input.ExpireDate,
		Guarantees:   input.Guarantees,
	}
	_, err = storage.CreateTokenDb(db.Repo, inputDB)
	if err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func GetToken(c *gin.Context, UriToken string, ContractId string) {
	num, err := strconv.Atoi(UriToken)
	if err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	res, err := lumx.FetchTokenType(ContractId, num)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusNotAcceptable)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func GetAllToken(c *gin.Context, ContractId string) {
	res, err := lumx.FetchTokenTypes(ContractId)
	if err != nil {
		c.Set("Error", "Error in lumx Api")
		c.Status(http.StatusNotAcceptable)
		return
	}
	c.Set("Response", res)
	c.Status(http.StatusOK)
}

func DistributePayments(c *gin.Context, ContractId string) {
	tokens, err := storage.GetAllTokensDb(db.Repo, ContractId)
	if err != nil {
		c.Set("Error", "Erro ao buscar tokens")
		c.Status(http.StatusInternalServerError)
		return
	}

	for _, token := range tokens {
		paymentAmount := calculatePayment(token.UriNumber, token.ContractID)
		err := lumx.DistributePayment(token.ContractID, paymentAmount)
		if err != nil {
			c.Set("Error", "Erro na distribuição de pagamentos")
			c.Status(http.StatusInternalServerError)
			return
		}
	}
	c.Status(http.StatusOK)
}

func EarlyRedemption(c *gin.Context, ContractId string, redemptionAmount float64) {
	tokens, err := storage.GetAllTokensDb(db.Repo, ContractId)
	if err != nil {
		c.Set("Error", "Erro ao buscar tokens")
		c.Status(http.StatusInternalServerError)
		return
	}

	for _, token := range tokens {
		premiumAmount := redemptionAmount + (redemptionAmount * 0.05)
		err := lumx.RedeemToken(token.ContractID, token.UriNumber, premiumAmount)
		if err != nil {
			c.Set("Error", "Erro no resgate antecipado")
			c.Status(http.StatusInternalServerError)
			return
		}
	}
	c.Status(http.StatusOK)
}

func isTokenExpired(expireDate string) bool {
	expireTime, _ := time.Parse("2006-01-02", expireDate)
	return time.Now().After(expireTime)
}
func calculateTokenAmount(assetValue float64) int {
	return int(assetValue)
}
func calculatePayment(tokenAmount uint64, contractId string) float64 {
	totalAmount := 100000.00 // Valor total do ativo
	return float64(tokenAmount) / float64(totalAmount)
}
