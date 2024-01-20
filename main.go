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
	apiService := restql.GinAPIService{Engine: router}

	// Define an API object
	config, _ := restql.ReadConfig("restql.yml")
	apis := restql.GetAPIs(config)

	// Use GenerateAPI to add a handler to the Gin engine
	restql.GenerateAPIs(apis, &apiService)

	// Start the Gin server
	router.Run(":8080")
}
