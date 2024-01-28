package models

import (
	"errors"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Game struct {
	ID          uuid.UUID `gorm:"type:uuid;default:not null"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Desc        string    `gorm:"type:text"`
	Image       string    `gorm:"type:varchar(255)"`
	Genre       string    `gorm:"type:varchar(100)"`
	ReleaseDate string    `gorm:"type:date"`
	gorm.Model
}

func (g Game) GetAll() ([]*Game, error) {
	db := GetDB()
	var games []*Game
	if err := db.Find(&games).Error; err != nil {
		log.Fatal(err)
	}
	return games, nil
}

func (g Game) GetByID(id string) (*Game, error) {
	db := GetDB()
	var game Game
	result := db.First(&game, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("Game not found")
	} else if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println(game)
	return &game, nil
}

func (g Game) Create(game *Game) (*Game, error) {
	db := GetDB()
	game.ID = uuid.NewV4()
	game.UpdatedAt = time.Now()

	err := db.Create(&game).Error
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when trying to save data to database")
	}
	return game, nil
}

func (g Game) Delete(id string) error {
	db := GetDB()
	result := db.Delete(&Game{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		return errors.New("error when trying to delete " + id + " data from database")
	}
	return nil
}
