package db

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func Init() {
	//TODO: Decide which DB to use from config
	db = CreateSqliteDB()
	InitSchema()
	CheckEmptyDB()
}

func CheckEmptyDB() {
	dbc := GetDB()
	row := dbc.QueryRow("SELECT COUNT(*) FROM users")
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		log.Println("Empty DB, fill with dummy data")
		LoadDummyData()
	}

}

func LoadDummyData() {
	insertData, err := os.ReadFile("./scripts/insert.sql")
	if err != nil {
		log.Fatal(err)
	}
	dbc := GetDB()
	stmt := string(insertData)
	_, err = dbc.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}

func InitSchema() {
	sqlBytes, err := os.ReadFile("./scripts/init.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Convert to string
	sqlString := string(sqlBytes)
	dbc := GetDB()
	// Execute SQL commands
	_, err = dbc.Exec(sqlString)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
