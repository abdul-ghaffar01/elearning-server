package middlewares

import (
	"elearning-server/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the Authorization header of incoming requests.
//
// This middleware expects the header in the format:
//     Authorization: Bearer <token>
//
// Steps:
//  1. Extracts the Authorization header.
//  2. Ensures it starts with "Bearer ".
//  3. Extracts the token portion.
//  4. Calls utils.VerifyJWT() to validate the token.
//  5. If the token is missing, malformed, or invalid, the request is aborted
//     with a 401 Unauthorized response.
//  6. If the token is valid, the request continues to the next handler.
//
// Returns a gin.HandlerFunc that performs this authentication check.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Expect: "Bearer <token>"
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}

		token := strings.Split(authHeader, " ")[1]

		// Extract the userID from JWT
		userID, err := utils.ExtractDataFromJwt(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Store userID into Gin context
		c.Set("user_id", userID)

		c.Next()
	}
}

