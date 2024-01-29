package models

import (
	"fmt"
	"log"

	"github.com/mistadave/gt-app/db"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var dbc *gorm.DB

func Init() {
	//TODO: Decide which DB to use from config
	dbc = db.GORMCreateSqliteDB()
	log.Println("Migrating models")
	MigrateModels()
	CheckEmptyDB()
}

func MigrateModels() {
	err := dbc.AutoMigrate(&User{})
	if err != nil {
		log.Println("Error migrating User model")
		panic(err)
	}
	err = dbc.AutoMigrate(&Tag{})
	if err != nil {
		log.Println("Error migrating Tag model")
		panic(err)
	}
	err = dbc.AutoMigrate(&Link{})
	if err != nil {
		log.Println("Error migrating Link model")
		panic(err)
	}
	err = dbc.AutoMigrate(&LinkTag{})
	if err != nil {
		log.Println("Error migrating LinkTag model")
		panic(err)
	}
	err = dbc.AutoMigrate(&Game{})
	if err != nil {
		log.Println("Error migrating Game model")
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return dbc
}

func CheckEmptyDB() {
	row := dbc.Table("users").Select("count(*)").Row()
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		log.Println("Empty DB, fill with dummy data")
		FillDummyData()
	}

}

func FillDummyData() {
	// Create dummy users
	for i := 0; i < 10; i++ {
		gender := "male"
		if i%2 != 0 {
			gender = "female"
		}
		user := User{
			ID:       uuid.FromStringOrNil(fmt.Sprintf("f0000000-0000-0000-0000-00000000000%d", i)),
			Name:     fmt.Sprintf("User%d", i),
			Username: fmt.Sprintf("username%d", i),
			Email:    fmt.Sprintf("user%d@example.com", i),
			Password: "password",
			Gender:   gender,
			Active:   true,
		}
		dbc.Create(&user)
	}

	// Create dummy tags
	for i := 0; i < 10; i++ {
		tag := Tag{
			Name: fmt.Sprintf("Tag%d", i),
		}
		dbc.Create(&tag)
	}

	for i := 0; i < 10; i++ {
		linkTag := LinkTag{
			LinkID: uuid.FromStringOrNil(fmt.Sprintf("f0000000-0000-0000-0000-00000000000%d", i)),
			TagID:  uint(i),
		}
		dbc.Create(&linkTag)
	}

	// Create dummy links
	for i := 0; i < 10; i++ {
		link := Link{
			ID:   uuid.FromStringOrNil(fmt.Sprintf("f0000000-0000-0000-0000-00000000000%d", i)),
			Name: fmt.Sprintf("Link%d", i),
			URL:  fmt.Sprintf("http://example.com/link%d", i),
		}
		dbc.Create(&link)
	}

	// Create dummy games
	for i := 0; i < 10; i++ {
		game := Game{
			ID:   uuid.FromStringOrNil(fmt.Sprintf("f0000000-0000-0000-0000-00000000000%d", i)),
			Name: fmt.Sprintf("Game%d", i),
		}
		dbc.Create(&game)
	}
}

// func LoadDummyData() {
// 	insertData, err := os.ReadFile("./scripts/insert.sql")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	dbc := GetDB()
// 	stmt := string(insertData)
// 	_, err = dbc.Exec(stmt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func InitSchema() {
// 	sqlBytes, err := os.ReadFile("./scripts/init.sql")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Convert to string
// 	sqlString := string(sqlBytes)
// 	dbc := GetDB()
// 	// Execute SQL commands
// 	_, err = dbc.Exec(sqlString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
