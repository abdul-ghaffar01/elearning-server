package database

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
	"sort"
)

// LoadAndRunSchema reads all SQL files from the schema folder
// and executes them in order based on their filename prefixes (001_, 002_, ...).
func LoadAndRunSchema(folderPath string) {
	// Get all .sql files from the folder
	files, err := filepath.Glob(filepath.Join(folderPath, "*.sql"))
	if err != nil {
		log.Fatalf("❌ Failed to read schema folder: %v", err)
	}

	if len(files) == 0 {
		log.Println("⚠️ No SQL files found in schema folder")
		return
	}

	// Sort files alphabetically to respect prefix order (001_, 002_, ...)
	sort.Strings(files)

	// Execute each file
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("❌ Failed to read file %s: %v", file, err)
		}

		_, err = DB.Exec(string(content))
		if err != nil {
			log.Fatalf("❌ Failed to execute SQL from file %s: %v", file, err)
		}

		fmt.Printf("✅ Executed %s successfully\n", filepath.Base(file))
	}
}
