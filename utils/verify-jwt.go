package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)


// VerifyJWT verifies a JWT token string and returns true if valid.
//
// Steps:
//  1. Parse and validate the token's signature.
//  2. Check expiry ("exp" claim).
//  3. Ensure signing method is HMAC (HS256).
//
// Returns:
//  - true, nil    → token is valid
//  - false, error → invalid or expired token

func VerifyJWT(tokenString string) (bool, error) {
	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure token is signed using HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		// token is invalid or expired
		return false, err
	}

	// Check token validity
	if !token.Valid {
		return false, errors.New("invalid token")
	}

	// TODO: will be checked if the token is not marked as expired in DB (for logout functionality)

	return true, nil
}
