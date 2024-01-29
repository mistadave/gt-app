package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type Tag struct {
	ID   uint   `gorm:"type:int;auto_increment;primary_key"`
	Name string `gorm:"type:varchar(255);not null"`
	gorm.Model
}

func (t Tag) SearchBy(query string) (*[]Tag, error) {
	db := GetDB()
	var tags []Tag
	result := db.Where("name LIKE ?", "%"+query+"%").Find(&tags)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("tag not found")
	} else if result.Error != nil {
		log.Fatal(result.Error)
	}
	return &tags, nil
}

func (t Tag) GetAll() ([]*Tag, error) {
	db := GetDB()
	var tags []*Tag
	result := db.Find(&tags)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("tags not found")
	} else if result.Error != nil {
		log.Fatal(result.Error)
	}
	return tags, nil
}

func (t Tag) Create(tag *Tag) (*Tag, error) {
	db := GetDB()
	err := db.Create(&tag).Error
	if err != nil {
		log.Fatal(err)
	}
	return tag, nil
}
