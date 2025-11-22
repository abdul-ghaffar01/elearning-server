package main

import (
	"elearning-server/routes"

	"github.com/gin-gonic/gin"
)

func main(){

	// Creating a gin router
	router := gin.Default();


	// setting up all the routes
	routes.SetupRoutes(router);

	// Starting the server on port 4406
	router.Run(":4406")
}