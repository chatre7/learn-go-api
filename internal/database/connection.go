package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"learn-api/pkg/errors"

	_ "github.com/lib/pq"
)

// DB represents the database connection
var DB *sql.DB

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB() error {
	var err error
	
	// Get database connection details from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "learnapi")

	// Create connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open database connection
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Check if connection is successful
	err = DB.Ping()
	if err != nil {
		return errors.ErrDatabase
	}

	fmt.Println("Successfully connected to PostgreSQL database!")
	return nil
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}