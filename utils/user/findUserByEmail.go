package user

import (
	"database/sql"
	"elearning-server/database"
	"elearning-server/types"
	"elearning-server/utils"
)
// FindUserByEmail retrieves a user record using the provided email address.
//
// The function loads the corresponding SQL query from the database/queries
// directory via utils.LoadQuery, executes it with the given email parameter,
// and scans the result into a types.User structure.
//
// If no user exists with the specified email address, the function returns
// sql.ErrNoRows so the caller can explicitly handle the "not found" case.
//
// Parameters:
//   - email: The email address that uniquely identifies the user.
//
// Returns:
//   - *types.User: The user entity if found.
//   - error: sql.ErrNoRows if no matching user exists, or another error if the
//     query fails to load, execute, or scan.
func FindUserByEmail(email string) (*types.User, error) {
	query, err := utils.LoadQuery("users/find_user_by_email")
	if err != nil {
		return nil, err
	}

	row := database.DB.QueryRow(query, email)

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

