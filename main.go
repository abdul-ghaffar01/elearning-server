package main

import (
	"elearning-server/database"
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

	// connecting with database
	database.Connect()
	defer database.CloseDB() // closing the database when main function ends
	
	// Creating the database schema
	database.LoadAndRunSchema("./database/schema")



	// Creating a gin router
	router := gin.Default()

	// setting up all the routes
	routes.SetupRoutes(router)

	// Starting the server on port 4406
	router.Run(":4406")
}
