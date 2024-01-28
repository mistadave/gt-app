package models

import (
	"log"
	"time"

	"github.com/mistadave/gt-api/db"
)

type Tag struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// type Tag struct {
// 	gorm.Model
// 	Name string
// }

func (t Tag) SearchBy(query string) (*Tag, error) {
	db := db.GetDB()
	sql := "SELECT * FROM tags WHERE name like %" + query + "%"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	tag := Tag{}
	for rows.Next() {
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.UpdatedAt); err != nil {
			log.Fatal(err)
		}
	}
	return &tag, nil
}

func (t Tag) GetAll() ([]*Tag, error) {
	db := db.GetDB()
	sql := "SELECT * FROM tags"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tags []*Tag
	for rows.Next() {
		var tag Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		tags = append(tags, &tag)
	}
	return tags, nil
}

func (t Tag) Create(tag *Tag) (*Tag, error) {
	db := db.GetDB()
	sql := "INSERT INTO tags (name) VALUES (?)"
	_, err := db.Exec(sql, tag.Name)
	if err != nil {
		log.Fatal(err)
	}
	return tag, nil
}
