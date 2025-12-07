package controllers

import (
	"elearning-server/utils/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request payload for signup
type SignupRequest struct {
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Signup(c *gin.Context) {
	var body SignupRequest

	// Parse + validate JSON
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// Call your CreateNewUser() function
	newUser, err := user.CreateNewUser(
		body.FullName,
		body.Email,
		body.Password,
		"", // No profile picture at signup
	)

	if err != nil {
		// Handle duplicate email error (Postgres unique constraint)
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email already registered",
			})
			return
		}

		// Other errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    newUser,
	})
}
