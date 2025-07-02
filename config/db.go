package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using environment variables")
    }

    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL not set")
    }

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping to DB:", err)
	}
	log.Println("Successfully connected to DB")
	return db
}
