// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gogean/restql"
)

func main() {
	// Create a Gin engine instance
	router := gin.Default()

	// Create an instance of GinAPIService
	ginAPIService := restql.GinAPIService{Engine: router}

	// Define an API object
	apiConfig, _ := restql.GetAPIConfig("restql.yml")
	apis := restql.GetAPIs(apiConfig)

	// Use GenerateAPI to add a handler to the Gin engine
	restql.GenerateAPIs(apis, &ginAPIService)

	// Start the Gin server
	router.Run(":8080")
}
