package user

import (
	"database/sql"
	"elearning-server/types"
	"log"
)

func FindOrCreateUser(name, email, pictureURL string) (*types.PublicUser, error) {

	// 1. Try to find user by email
	u, err := FindUserByEmail(email)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// 2. If user exists
	if u != nil {
		log.Println("User already exists")

		// Update picture if needed
		if pictureURL != "" && u.Profile != pictureURL {
			updated, err := UpdateUserPicture(email, pictureURL)
			if err != nil {
				return nil, err
			}
			return updated, nil
		}

		// Convert to public user
		public := types.ToPublicUser(u)
		return &public, nil
	}

	// 3. If user does NOT exist → Create
	log.Println("Creating new user")
	newUser, err := CreateNewUser(name, email, "", pictureURL)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
