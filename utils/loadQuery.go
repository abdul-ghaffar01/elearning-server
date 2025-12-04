package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// LoadQuery reads an SQL file from database/query folder and returns its contents.
func LoadQuery(filename string) (string, error) {
	// Build path: project_root/database/query/filename
	path := filepath.Join("database", "queries", filename + ".sql")

	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to load query file %s: %w", filename, err)
	}

	return string(content), nil
}
