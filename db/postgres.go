package db

import (
	"database/sql"
	"log"
)

func CreatePsqlDB() *sql.DB {
	//TODO: Use postgresql settings from config
	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Create a table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
