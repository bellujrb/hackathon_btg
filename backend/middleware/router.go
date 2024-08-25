package middleware

import (
	_ "brasa/docs"
	receivable "brasa/internal/receivable/handler"
	tokens "brasa/internal/tokens/handler"
	user "brasa/internal/user/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Brasa
// @version 0.1
// @description API
// @termsOfService http://swagger.io/terms/
// @host 52.15.128.117:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(ResponseHandler())

	//Use response, but not Token
	r.GET("/token", generateTokenHandler)
	r.PUT("/login", LoginUser)
	r.POST("/create-user", user.CreateUser)

	auth := r.Group("/api")

	auth.Use(authMiddleware)

	//Response and token service

	auth.POST("/create-contract", user.CreateContract)

	auth.POST("/deploy-contract", user.DeployContract)
	auth.GET("/contract-details", user.PullContract)
	auth.GET("/all-contract", user.PullAllContract)

	auth.PUT("/logged")
	auth.GET("/user", user.PullUser)

	auth.POST("/create-token", tokens.CreateToken)
	auth.GET("/get-token", tokens.GetToken)
	auth.GET("/all-token", tokens.GetAllToken)

	auth.POST("/receivables", receivable.CreateReceivable)
	auth.GET("/receivables/:id", receivable.GetReceivableByID)
	auth.GET("/receivables", receivable.GetAllReceivables)

	return r
}
