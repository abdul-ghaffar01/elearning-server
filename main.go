package main

import (
	"elearning-server/routes"
	
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Creating a gin router
	router := gin.Default()

	// setting up all the routes
	routes.SetupRoutes(router)

	// Starting the server on port 4406
	router.Run(":4406")
}
