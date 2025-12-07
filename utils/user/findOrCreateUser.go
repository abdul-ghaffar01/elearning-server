package user

import (
	"database/sql"
	"log"
	"elearning-server/types"
)

func FindOrCreateUser(name, email, pictureURL string) (*types.User, error) {

	// 1. Try to find user by email
	user, err := FindUserByEmail(email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// 2. If user exists:
	if user != nil {
		log.Println("User already exists")

		// If pictureURL provided → update
		if pictureURL != "" && user.Profile != pictureURL {
			return UpdateUserPicture(email, pictureURL)
		}

		return user, nil
	}

	// 3. If user does not exist → create one
	log.Println("Creating new user")
	return CreateNewUser(name, email, "", pictureURL)
}

