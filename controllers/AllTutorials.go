package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllTutorials(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tutorials": "All tutorials",
	})
}