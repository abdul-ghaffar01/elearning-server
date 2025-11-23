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

	optionalAuth := router.Group("")
	optionalAuth.Use(middlewares.OptionalAuthMiddleware())
	{
		optionalAuth.GET("/tutorials", controllers.AllTutorials)
		optionalAuth.GET("/tutorials/:id", controllers.SingleTutorial)
	}

	// Public routes (no prefix)
	public := router.Group("")
	{
		public.GET("/google-login", controllers.GoogleLogin)
		public.GET("/google/callback", controllers.GoogleCallback)
	}
}
