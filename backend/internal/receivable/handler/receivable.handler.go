package receivable

import (
	"brasa/internal/receivable/service"

	"github.com/gin-gonic/gin"
)

// CreateReceivable godoc
// @Summary Cria um novo recebível
// @Description Cria um novo recebível a partir dos dados fornecidos
// @Tags Receivable
// @Accept json
// @Produce json
// @Param receivable body db.Receivable true "Dados do Recebível"
// @Success 200 {object} db.Receivable
// @Failure 400 {object} gin.H
// @Router /receivables [post]
func CreateReceivable(c *gin.Context) {
	service.CreateReceivable(c)
}

// GetReceivableByID godoc
// @Summary Obtém um recebível pelo ID
// @Description Busca um recebível pelo ID fornecido
// @Tags Receivable
// @Accept json
// @Produce json
// @Param id path string true "ID do Recebível"
// @Success 200 {object} db.Receivable
// @Failure 404 {object} gin.H
// @Router /receivables/{id} [get]
func GetReceivableByID(c *gin.Context) {
	service.GetReceivableByID(c)
}

// GetAllReceivables godoc
// @Summary Obtém todos os recebíveis
// @Description Retorna todos os recebíveis cadastrados
// @Tags Receivable
// @Accept json
// @Produce json
// @Success 200 {array} db.Receivable
// @Failure 500 {object} gin.H
// @Router /receivables [get]
func GetAllReceivables(c *gin.Context) {
	service.GetAllReceivables(c)
}
