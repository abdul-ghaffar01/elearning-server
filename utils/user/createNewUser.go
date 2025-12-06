package user

import (
	"elearning-server/database"
	"elearning-server/types"
	"elearning-server/utils"
)

func CreateNewUser(name, email, pictureURL string) (*types.User, error) {
	query, err := utils.LoadQuery("users/create_new_user")
	if err != nil {
		return nil, err
	}
	row := database.DB.QueryRow(query, name, email, pictureURL)

	var u types.User
	err = row.Scan(&u.ID, &u.FullName, &u.Email, &u.Profile, &u.Joined, &u.ProfileSetupped)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
