package database

import (
	"log"
	"os"
	"path/filepath"
)


// LoadQuery reads a SQL file from the given path and returns it as a string.
// Example usage:
//    query := LoadQuery("./database/queries/tutorials.sql")
func LoadQuery(path string) string {
	// Clean the path
	fullPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("❌ Failed to get absolute path for %s: %v", path, err)
	}

	// Read file contents
	data, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("❌ Failed to read SQL file %s: %v", fullPath, err)
	}

	return string(data)
}
