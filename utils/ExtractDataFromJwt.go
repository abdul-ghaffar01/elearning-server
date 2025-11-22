package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)


// ExtractDataFromJwt parses a JWT token and returns the "user_id" claim.
//
// Steps:
//  1. Parse token with the secret key.
//  2. Validate signature + check expiration.
//  3. Extract "user_id" from claims.
// 
// Returns:
//  - userID string and nil if successful
//  - empty string and error if invalid or missing user_id
func ExtractDataFromJwt(tokenString string) (string, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract claims as map
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	// Get the user_id claim
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id not found in token")
	}

	return userID, nil
}
