package controllers

import (
	"elearning-server/utils/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserDetails handles GET /user requests.
//
// It extracts the "id" from access token provided in headers.
// Then, it fetches the user details from the database and returns them in the response.
func UserDetails(c *gin.Context) {
		// 1. Get userId from context
	userIDRaw, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	// 2. Type assert to uuid.UUID
	userId, ok := userIDRaw.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id type",
		})
		return
	}
	// 3. Fetch user from database
	u, err := user.FindUserById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	// 4. Return user
	c.JSON(http.StatusOK, u)
}

