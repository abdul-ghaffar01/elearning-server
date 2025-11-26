package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Typically you'll load env vars
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE") // usually "disable" on local
	if ssl == "" {
		ssl = "disable"
	}

	// Build connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, ssl,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Failed to build connection:", err)
	}

	// Check connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Cannot connect to PostgreSQL:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
}

func CloseDB() {
	log.Println("Closing db")
	if DB != nil {
		DB.Close()
		log.Println("🔌 Database connection closed.")
	}
}
