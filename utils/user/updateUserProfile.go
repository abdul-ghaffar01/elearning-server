package user

import (
	"elearning-server/database"
	"elearning-server/types"
	"elearning-server/utils"
)

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
