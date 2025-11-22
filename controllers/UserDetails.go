package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// UserDetails handles GET /user/:id requests.
//
// It extracts the "id" parameter from the URL and returns it as JSON.
// This is typically used to fetch details for a specific user.
func UserDetails(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
