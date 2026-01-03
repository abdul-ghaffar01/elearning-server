package user

import (
	"database/sql"
	"elearning-server/database"
	"elearning-server/types"
	"elearning-server/utils"

	"github.com/google/uuid"
)
// FindUserById retrieves a user record by its unique identifier.
//
// The function loads the SQL query from the database/queries directory using
// utils.LoadQuery, executes it with the provided UUID, and maps the result to a
// types.User structure.
//
// If no user exists for the given ID, the function returns sql.ErrNoRows,
// allowing callers to explicitly handle the "not found" case.
//
// Parameters:
//   - userId: The UUID that uniquely identifies the user.
//
// Returns:
//   - *types.User: The user entity if found.
//   - error: sql.ErrNoRows if no matching user exists, or another error if the
//     query fails to load, execute, or scan.
func FindUserById(userId uuid.UUID) (*types.User, error) {
	query, err := utils.LoadQuery("users/find_user_by_id")
	if err != nil {
		return nil, err
	}

	row := database.DB.QueryRow(query, userId)

	var u types.User
	err = row.Scan(
		&u.ID,
		&u.FullName,
		&u.Email,
		&u.Profile,
		&u.Joined,
		&u.ProfileSetupped,
	)

	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}
