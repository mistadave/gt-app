package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GORMCreateSqliteDB() *gorm.DB {
	log.Println("CreateSqliteDB")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database")
		log.Fatal(err)
	}
	return db
}

func CreateSqliteDB() *sql.DB {
	log.Println("CreateSqliteDB")
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("Error opening database")
		log.Fatal(err)
	}
	return db
}
