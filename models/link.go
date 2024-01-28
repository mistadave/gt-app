package models

import (
	"log"
	"time"

	"github.com/mistadave/gt-api/db"
	uuid "github.com/satori/go.uuid"
)

type LinkTag struct {
	LinkID string `db:"link_id"`
	TagID  int    `db:"tag_id"`
}

type Link struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Desc      string    `db:"desc" json:"desc"`
	URL       string    `db:"url" json:"url"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
	GameID    uuid.UUID `db:"game_id" json:"gameId"`
}

func (l Link) GetAll() ([]*Link, error) {
	db := db.GetDB()
	sql := "SELECT * FROM links"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var links []*Link
	for rows.Next() {
		var link Link
		if err := rows.Scan(&link.ID, &link.Name, &link.Desc, &link.URL, &link.UpdatedAt, &link.GameID); err != nil {
			log.Fatal(err)
		}
		links = append(links, &link)
	}
	return links, nil
}

func (l Link) GetByID(id string) (*Link, error) {
	db := db.GetDB()
	sql := "SELECT * FROM links WHERE id = " + id
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	link := Link{}
	for rows.Next() {
		if err := rows.Scan(&link.ID, &link.Name, &link.Desc, &link.URL, &link.UpdatedAt, link.GameID); err != nil {
			log.Fatal(err)
		}
	}
	return &link, nil
}

func (l Link) GetByGameID(id string) (*[]Link, error) {
	db := db.GetDB()
	sql := "SELECT * FROM links WHERE game_id = ? limit 10"
	rows, err := db.Query(sql, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	for rows.Next() {
		link := Link{}
		if err := rows.Scan(&link.ID, &link.Name, &link.Desc, &link.URL, &link.UpdatedAt, &link.GameID); err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	return &links, nil
}

func (l Link) Create(link *Link) (*Link, error) {
	db := db.GetDB()
	id := uuid.NewV4()
	link.ID = id
	link.UpdatedAt = time.Now()
	stmt := `INSERT INTO links (id, name, desc, url, updated_at, game_id) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(stmt, link.ID, link.Name, link.Desc, link.URL, link.UpdatedAt, link.GameID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return link, nil
}

func (l Link) Delete(id string) error {
	db := db.GetDB()
	sql := "DELETE FROM links WHERE id = ?"
	_, err := db.Exec(sql, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
