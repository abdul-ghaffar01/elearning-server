package user

import (
	"elearning-server/database"
	"elearning-server/types"
	"elearning-server/utils"
)
// UpdateUserPicture updates the profile picture URL of a user identified by email
// and returns the updated public user record.
//
// The function loads the corresponding SQL update query from the
// database/queries directory using utils.LoadQuery, executes it with the provided
// email and pictureURL parameters, and scans the updated user data into a
// types.PublicUser structure.
//
// The underlying SQL query is expected to return the updated user fields
// (ID, full name, email, profile picture URL, and joined date) as part of the
// update operation.
//
// Parameters:
//   - email: The email address that uniquely identifies the user.
//   - pictureURL: The new profile picture URL to be stored.
//
// Returns:
//   - *types.PublicUser: The updated public representation of the user.
//   - error: A non-nil error if the query cannot be loaded, executed, or scanned.
func UpdateUserPicture(email, pictureURL string) (*types.PublicUser, error) {
	query, err := utils.LoadQuery("users/update_user_profile")
	if err != nil {
		return nil, err
	}

	row := database.DB.QueryRow(query, email, pictureURL)

	var u types.PublicUser
	err = row.Scan(&u.ID, &u.FullName, &u.Email, &u.Profile, &u.Joined)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
