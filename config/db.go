package config

import (
	"database/sql"
	"log"
)

func ConnectDB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=admin1234 dbname=tasksdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping to DB:", err)
	}
	log.Println("Successfully connected to DB")
	return db
}
