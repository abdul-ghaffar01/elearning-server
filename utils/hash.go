package utils

import (
	"golang.org/x/crypto/bcrypt"
)


// HashPassword generates a bcrypt hash from a plain-text password.
//
// Parameters:
//   - password (string): The raw password to hash.
//
// Returns:
//   - string: The hashed password to store in the database.
//   - error: Any hashing error.
//
// Usage:
//   hashed, err := utils.HashPassword("mypassword")
//
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}


// CheckPasswordHash compares a plain-text password with a bcrypt hashed password.
//
// Parameters:
//   - password (string): Raw password submitted by user.
//   - hash (string): Password hash stored in database.
//
// Returns:
//   - bool: true if the password matches, false otherwise.
//
// Security:
//   - bcrypt automatically handles salting and secure comparison.
//   - Timing-safe comparison prevents timing attacks.
//
// Usage:
//   isValid := utils.CheckPasswordHash("input", user.PasswordHash)
//
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
