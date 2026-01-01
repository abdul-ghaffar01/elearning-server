package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// LoadQuery loads an SQL query from the database/queries directory and returns
// its contents as a string.
//
// The function constructs the file path using the provided filename by appending
// the ".sql" extension and joining it with the relative path
// "database/queries". For example, passing "get_users" will attempt to read
// "database/queries/get_users.sql".
//
// If the file cannot be read (due to reasons such as a missing file or
// insufficient permissions), an error is returned with contextual information.
//
// Parameters:
//   - filename: The base name of the SQL file (without the ".sql" extension).
//
// Returns:
//   - string: The contents of the SQL file.
//   - error: A non-nil error if the file could not be read.
func LoadQuery(filename string) (string, error) {
	// Build path: project_root/database/queries/filename.sql
	path := filepath.Join("database", "queries", filename+".sql")

	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to load query file %s: %w", filename, err)
	}

	return string(content), nil
}
