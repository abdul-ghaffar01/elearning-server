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

type PublicUser struct {
	ID             uuid.UUID `json:"id"`               // Primary Key (UUID)
	Email          string    `json:"email"`            // User email
	FullName       string    `json:"full_name"`        // Full name
	DateOfBirth    string    `json:"date_of_birth"`    // DOB as YYYY-MM-DD string
	Country        string    `json:"country"`          // Country of residence
	Role           string    `json:"role"`             // Role (e.g., admin, user)
	Profile        string    `json:"profile"`          // Profile description or JSON
	Joined         time.Time `json:"joined"`           // Account creation date
	Deactivated    bool      `json:"deactivated"`      // Is account deactivated
	ProfileSetupped bool     `json:"profile_setupped"`  // Profile setup completed
}

func ToPublicUser(u *User) PublicUser {
	var dob string
	if u.DateOfBirth.Valid {
		dob = u.DateOfBirth.Time.Format("2006-01-02")
	}

	return PublicUser{
		ID:              u.ID,
		Email:           u.Email,
		FullName:        u.FullName,
		DateOfBirth:     dob,
		Country:         u.Country,
		Role:            u.Role,
		Profile:         u.Profile,
		Joined:          u.Joined,
		Deactivated:     u.Deactivated,
		ProfileSetupped: u.ProfileSetupped,
	}
}
