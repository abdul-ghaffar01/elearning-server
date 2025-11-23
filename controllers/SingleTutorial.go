package controllers

import "github.com/gin-gonic/gin"

func SingleTutorial(c *gin.Context) {
	// Implementation for fetching and returning a single tutorial by ID
	id := c.Param("id")

	c.JSON(200, gin.H{
		"id": id,
	})
}
