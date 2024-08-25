package main

import (
	"brasa/db"
	"brasa/middleware"
)

// @title           Brasa
// @version         1.0
// @description     This is a server for app.

// @host      52.15.128.117:8080

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := middleware.SetupRouter()
	db.ConnectDatabase()
	r.Run(":8080")
}
