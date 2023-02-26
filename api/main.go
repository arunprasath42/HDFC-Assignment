// Main function - entry point of the application - parses the command line arguments and executes the corresponding command
package main

import (
	"hdfc-assignment/config"
	"hdfc-assignment/route"
	"hdfc-assignment/utils/validator"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	validator.Init()
	router := gin.Default()
	route.SetupRoutes(router)
}
