package main

import (
	"log"
	"time"

	"elearning-server/database"
	"elearning-server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect database
	database.Connect()
	defer database.CloseDB()

	// Load schema
	database.LoadAndRunSchema("./database/schema")

	// Create router
	router := gin.Default()

	// ✅ CORS CONFIG (MUST BE BEFORE ROUTES)
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://elearning.iabdulghaffar.com",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Authorization",
		},
		ExposeHeaders: []string{
			"Set-Cookie",
		},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	router.Run(":4406")
}
