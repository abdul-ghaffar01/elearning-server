package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Secret key for signing tokens.
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT generates a signed JWT token for a given userID.
//
// The token contains:
// - "user_id" (uuid) : the ID of the user
// - "exp" (int64)      : expiration time (15 minutes for access, 7 days for refresh)
//
// Returns the token string or an error.

func GenerateJWT(userID uuid.UUID, typeOfToken string) (string, error) {
	// Create token claims
	var claims jwt.MapClaims
	if typeOfToken == "refresh" {
		claims = jwt.MapClaims{
			"user_id": userID,
			"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
		}
	} else {
		claims = jwt.MapClaims{
			"user_id": userID,
			"exp":     time.Now().Add(time.Minute * 15).Unix(),
		}
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}


	return tokenString, nil
}
