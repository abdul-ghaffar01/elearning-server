package controllers

import "github.com/gin-gonic/gin"


func Signup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Signup endpoint - to be implemented",
	})
}