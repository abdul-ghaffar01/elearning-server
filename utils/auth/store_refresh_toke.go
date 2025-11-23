package auth

import "elearning-server/database"

func StoreRefreshToken(userID string, refreshToken string) error {
	// Implement the logic to store the refresh token in the database
	// This is a placeholder implementation and should be replaced with actual database code

	// TODO: Store refresh tokens in DB (not implemented here)

	query := database.LoadQuery("./database/queries/store_refresh_token.sql")
	_, err := database.DB.Exec(query, userID, refreshToken)
	if err != nil {
		return err
	}

	return nil
}
