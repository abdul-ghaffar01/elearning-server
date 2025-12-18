package user

import (
	"database/sql"
	"elearning-server/database"
	"elearning-server/types"
	"elearning-server/utils"

	"github.com/google/uuid"
)

func FindUserById(userId uuid.UUID) (*types.User, error) {
	query, err := utils.LoadQuery("users/find_user_by_id")
	if err != nil {
		return nil, err
	}
	row := database.DB.QueryRow(query, userId)

	var u types.User
	err = row.Scan(&u.ID, &u.FullName, &u.Email, &u.Profile, &u.Joined, &u.ProfileSetupped)

	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}
