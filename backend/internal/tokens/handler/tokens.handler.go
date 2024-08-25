package tokens

import (
	inter "brasa/internal/tokens/interface"
	"brasa/internal/tokens/service"
	errors "brasa/middleware/interfaces/errors"

	"github.com/gin-gonic/gin"
)

// @Summary Make Tokens
// @Description Create a new user in db
// @Tags Token
// @Accept json
// @Produce json
// @Param request body inter.TokenCreateInput true "Data for make a new token"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.TokenCreateOutput "New Token Created successfully"
// @Failure 500 {object} errors.InternalServerError "Unable to created a new token"
// @Router /api/create-token [post]
func CreateToken(c *gin.Context) {
	var input inter.TokenCreateInput
	if err := c.BindJSON(&input); err != nil {
		c.Set("Response", "Invalid parameters, need a JSON with email")
		c.Status(errors.StatusNotAcceptable)
		return
	}
	service.CreateToken(c, input)
}

// @Summary Get specify token
// @Description Create a new user in db
// @Tags Token
// @Accept json
// @Produce json
// @Param ContractId header string true "E-mail from user"
// @Param UriToken header string true "Its a uri content in tokens"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.UserOutputController "New User Created successfully"
// @Failure 500 {object} errors.InternalServerError "Unable to store data in database"
// @Router /api/get-token [get]
func GetToken(c *gin.Context) {
	ContractId := c.GetHeader("ContractId")
	UriToken := c.GetHeader("UriToken")
	service.GetToken(c, ContractId, UriToken)
}

// @Summary Get all Tokens
// @Description Create a new user in db
// @Tags Token
// @Accept json
// @Produce json
// @Param ContractId header string true "contract id"
// @Param Authorization header string true "Token de autenticação (Colocar o token deixando o Bearer)" default(Bearer <token>)
// @Success 200 {object} inter.AllTokensOutput "New User Created successfully"
// @Failure 500 {object} errors.InternalServerError "Unable to store data in database"
// @Router /api/all-token [get]
func GetAllToken(c *gin.Context) {
	ContractId := c.GetHeader("ContractId")
	service.GetAllToken(c, ContractId)
}
