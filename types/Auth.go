package types

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID           uuid.UUID     `json:"id"`             // Primary Key (UUID)
	UserID       uuid.UUID     `json:"user_id"`        // Foreign Key referencing UserProfile
	RefreshToken string        `json:"refresh_token"`  // JWT refresh token
	IsExpired    bool          `json:"is_expired"`     // Session expired or not
	IPAddress    string        `json:"ip_address"`     // Client IP address
	LoginTime    time.Time     `json:"login_time"`     // Login timestamp
	LogoutTime   sql.NullTime  `json:"logout_time"`    // Logout timestamp, nullable
	DeviceType   string        `json:"device_type"`    // Desktop / Mobile / Tablet
	OS           string        `json:"os"`             // Operating system
	Browser      string        `json:"browser"`        // Browser name
	Country      string        `json:"country"`        // Geo info
	City         string        `json:"city"`           // Geo info
}
