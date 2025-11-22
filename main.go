package main

import (
	"elearning-server/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"elearning-server/utils"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Creating a gin router
	router := gin.Default()

	token, err := utils.GenerateJWT("user123", "access")
	log.Println("Generated Token:", token, "Error:", err)

	verfication1, err := utils.VerifyJWT(token)
	log.Println("Token Verification1:", verfication1, "Error:", err)

	verification, err := utils.VerifyJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjM4MTg4MTAsInVzZXJfaWQiOiJ1c2VyMTIzIn0.A9EfUASmiBSyFdysiG7Lsw7RHtCkr_ogcu9f0DAsph8")
	log.Println("Token Verification:", verification, "Error:", err)	

	// setting up all the routes
	routes.SetupRoutes(router)

	// Starting the server on port 4406
	router.Run(":4406")
}
