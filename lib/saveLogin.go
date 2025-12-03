package lib

import (
	"elearning-server/database"
	"elearning-server/utils"
	"log"

	"github.com/google/uuid"
)

// SaveLogin inserts a login record for a user
func SaveLogin(userId uuid.UUID,
	refreshToken,
	ip,
	deviceType,
	os,
	browser,
	country,
	city string) error {

	query, err := utils.LoadQuery("auth/save_login")
	if err != nil {
		log.Println("❌ Failed to load query for saving login:", err)
		return err
	}

	_, err = database.DB.Exec(
		query,
		userId,
		refreshToken,
		ip,
		deviceType,
		os,
		browser,
		country,
		city,
	)

	if err != nil {
		log.Println("❌ Failed to save login:", err)
		return err
	}

	return nil
}
