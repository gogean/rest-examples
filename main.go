// main.go

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gogean/restql"
)

func main() {
	dbConnectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s", os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PWD"),
		os.Getenv("APP_DB_NAME"))

	db, _ := restql.SQLConnect(dbConnectionString)
	dbConnection := restql.SQL{Connection: db}

	// Create a Gin engine instance
	router := gin.Default()

	// Create an instance of GinAPIService
	ginAPIService := restql.GinAPIService{Engine: router}

	// Define an api object
	apiConfig, _ := restql.GetAPIConfig("../restql/test-config/restql.yml")
	apis := restql.GetAPIs(apiConfig)

	// Use generateAPI to add a handler to the Gin engine
	restql.GenerateAPIs(apis, &ginAPIService, &dbConnection)

	// Start the Gin server
	router.Run(":8080")
}
