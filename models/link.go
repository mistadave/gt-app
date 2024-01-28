package models

import (
	"errors"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type LinkTag struct {
	LinkID uuid.UUID `gorm:"type:uuid;not null"`
	TagID  uint      `gorm:"type:int;not null"`
	gorm.Model
}

type Link struct {
	ID     uuid.UUID `gorm:"type:uuid;default:not null"`
	Name   string    `gorm:"type:varchar(100);not null"`
	Desc   string    `gorm:"type:text"`
	URL    string    `gorm:"type:varchar(255);not null"`
	GameID uuid.UUID `gorm:"type:uuid;not null"`
	gorm.Model
}

func (l Link) GetAll() ([]*Link, error) {
	db := GetDB()
	var links []*Link
	result := db.Find(&links)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("Link not found")
	} else if result.Error != nil {
		log.Fatal(result.Error)
	}
	return links, nil
}

func (l Link) GetByID(id string) (*Link, error) {
	db := GetDB()
	link := &Link{}
	result := db.Where("id = ?", id).First(link)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("Link not found")
	} else if result.Error != nil {
		log.Fatal(result.Error)
	}
	return link, nil
}

func (l Link) GetByGameID(id string) ([]Link, error) {
	db := GetDB()
	var links []Link
	result := db.Where("game_id = ?", id).Limit(10).Find(&links)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("Links not found")
	} else if result.Error != nil {
		log.Fatal(result.Error)
	}
	return links, nil
}

func (l Link) Create(link *Link) (*Link, error) {
	db := GetDB()
	link.ID = uuid.NewV4()
	link.UpdatedAt = time.Now()
	result := db.Create(link)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return link, nil
}

func (l Link) Delete(id string) error {
	db := GetDB()
	link := &Link{}
	result := db.Where("id = ?", id).Delete(link)
	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	return nil
}
