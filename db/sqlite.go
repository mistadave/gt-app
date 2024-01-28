package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GORMCreateSqliteDB() *gorm.DB {
	log.Println("CreateSqliteDB")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database")
		log.Fatal(err)
		panic("failed to connect database")
	}
	return db
}
