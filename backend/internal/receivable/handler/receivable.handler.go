package receivable

import (
	"brasa/internal/receivable/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateReceivable godoc
// @Summary Creates a new receivable
// @Description Creates a new receivable from the provided data
// @Tags Receivable
// @Accept json
// @Produce json
// @Param receivable body db.Receivable true "Receivable data"
// @Success 200 {object} db.Receivable
// @Failure 400 {object} ErrorResponse
// @Router /receivables [post]
func CreateReceivable(c *gin.Context) {
	service.CreateReceivable(c)
}

// GetReceivableByID godoc
// @Summary Get a receivable by ID
// @Description Fetches a receivable by its ID
// @Tags Receivable
// @Accept json
// @Produce json
// @Param id path string true "Receivable ID"
// @Success 200 {object} db.Receivable
// @Failure 404 {object} ErrorResponse
// @Router /receivables/{id} [get]
func GetReceivableByID(c *gin.Context) {
	service.GetReceivableByID(c)
}

// GetAllReceivables godoc
// @Summary Get all receivables
// @Description Returns all registered receivables
// @Tags Receivable
// @Accept json
// @Produce json
// @Success 200 {array} db.Receivable
// @Failure 500 {object} ErrorResponse
// @Router /receivables [get]
func GetAllReceivables(c *gin.Context) {
	service.GetAllReceivables(c)
}

// UpdatePaymentDate godoc
// @Summary Updates the payment date of a receivable
// @Description Changes the payment date of the given receivable
// @Tags Receivable
// @Accept json
// @Produce json
// @Param id path string true "Receivable ID"
// @Param payment_date body string true "New Payment Date"
// @Success 200 {object} db.Receivable
// @Failure 500 {object} ErrorResponse
// @Router /receivables/{id}/payment_date [put]
func UpdatePaymentDate(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		PaymentDate string `json:"payment_date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdatePaymentDate(id, input.PaymentDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating payment date"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment date updated successfully"})
}

// GetEarlyPaymentOptions godoc
// @Summary Fetches early payment options for a receivable
// @Description Returns options for anticipating the receivable's payment
// @Tags Receivable
// @Accept json
// @Produce json
// @Param id path string true "Receivable ID"
// @Success 200 {object} db.Receivable
// @Failure 500 {object} ErrorResponse
// @Router /receivables/{id}/early_payment [get]
func GetEarlyPaymentOptions(c *gin.Context) {
	id := c.Param("id")

	options, err := service.GetEarlyPaymentOptions(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Options not found"})
		return
	}

	c.JSON(http.StatusOK, options)
}
