package routes

import (
	"elearning-server/controllers"
	"elearning-server/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.PongController)

	// Auth required routes (no prefix)
	auth := router.Group("")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/dashboard", controllers.DashboardController)

		// User related routes
		auth.GET("/user/:id", controllers.UserDetails)
	}

	// Public routes (no prefix)
	public := router.Group("")
	{
		public.GET("/tutorials", controllers.AllTutorials)
	}
}
