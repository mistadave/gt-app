package models

import (
	"log"
)

type Genre struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"type:varchar(100);not null"`
}

func (g Genre) GetAll() ([]*Genre, error) {
	db := GetDB()
	var genres []*Genre
	if err := db.Find(&genres).Error; err != nil {
		log.Fatal(err)
	}
	return genres, nil
}
func (g Genre) GetByID(id string) (*Genre, error) {
	db := GetDB()
	var genre Genre
	result := db.First(&genre, "id = ?", id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return &genre, nil
}
func (g Genre) Create(genre *Genre) (*Genre, error) {
	db := GetDB()
	err := db.Create(&genre).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return genre, nil
}

func (g Genre) Update(genre *Genre) (*Genre, error) {
	db := GetDB()
	err := db.Save(&genre).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return genre, nil
}

func (g Genre) Delete(id string) error {
	db := GetDB()
	result := db.Delete(&Genre{}, "id = ?", id)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
