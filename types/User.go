package types

import (
	"database/sql"
	"time"
	"github.com/google/uuid"
)


type User struct {
	ID             uuid.UUID     `json:"id"`               // Primary Key (UUID)
	Email          string        `json:"email"`            // User email
	FullName       string        `json:"full_name"`        // Full name
	DateOfBirth    sql.NullTime  `json:"date_of_birth"`    // Nullable DOB
	Country        string        `json:"country"`          // Country of residence
	Role           string        `json:"role"`             // Role (e.g., admin, user)
	Profile        string        `json:"profile"`          // Profile description or JSON
	Joined         time.Time     `json:"joined"`           // Account creation date
	Password       string        `json:"password"`         // Hashed password
	Deactivated    bool          `json:"deactivated"`      // Is account deactivated
	ProfileSetupped bool         `json:"profile_setupped"`  // Profile setup completed
	UpdatedAt      time.Time     `json:"updated_at"`       // Last updated timestamp
}