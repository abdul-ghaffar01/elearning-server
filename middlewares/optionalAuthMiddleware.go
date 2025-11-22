package middlewares

import (
	"elearning-server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// OptionalAuthMiddleware attempts to authenticate the user, 
// but does NOT block the request if authentication fails.
//
// Behavior:
//  1. Checks for Authorization header ("Bearer <token>").
//  2. If the header is missing/invalid → user is treated as *not logged in*.
//  3. If a token exists → it is validated using utils.VerifyJWT().
//  4. If valid → userID is stored in Gin context: c.Set("userID", <value>).
//  5. Request continues regardless of login state.
//
// This middleware is useful for routes that should work for both:
//  - Public users (no login)
//  - Logged-in users (personalized content)


func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// If no token → proceed as guest
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.Next()
			return
		}

		// Extract the token
		token := strings.Split(authHeader, " ")[1]

		// Verify token
		userID, err := utils.VerifyJWT(token)
		if err != nil {
			// Token invalid → still allow request but as guest
			c.Next()
			return
		}

		// Save userID for controller access
		c.Set("userID", userID)

		c.Next()
	}
}
