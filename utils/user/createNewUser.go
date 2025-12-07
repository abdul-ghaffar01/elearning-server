package user

import (
	"elearning-server/database"
	"elearning-server/types"
	"elearning-server/utils"
	"errors"
)

/*
CreateUser creates a new user in the database using flexible input fields.
Required fields:
  - FullName
  - Email

Optional fields:
  - Password (will be hashed automatically)
  - Picture (stored in 'profile')

This function:
  - Makes optional fields empty
  - Automatically hashes passwords if provided
  - Loads SQL dynamically from queries/users/create_new_user.sql
  - Scans all returned user fields into types.User
*/
func CreateNewUser(fullname, email, password, profile string) (*types.PublicUser, error) {

	if fullname == "" || email == "" {
		return nil, errors.New("fullname and email are required")
	}

	// Hash password if provided
	var hashedPassword *string
	if password != "" {
		h, err := utils.HashPassword(password)
		if err != nil {
			return nil, err
		}
		hashedPassword = &h
	}

	// Load SQL
	query, err := utils.LoadQuery("users/create_new_user")
	if err != nil {
		return nil, err
	}

	// Execute query
	row := database.DB.QueryRow(
		query,
		fullname,
		email,
		hashedPassword,
		profile,
	)

	// Scan user
	var u types.PublicUser
	err = row.Scan(
		&u.ID,
		&u.FullName,
		&u.Email,
		&u.Profile,
		&u.ProfileSetupped,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

