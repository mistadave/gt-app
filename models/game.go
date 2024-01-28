package models

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/mistadave/gt-api/db"
	"github.com/mistadave/gt-api/form"
	uuid "github.com/satori/go.uuid"
)

type Game struct {
	ID          uuid.UUID `db:"id,omitempty"`
	Name        string    `db:"name"`
	Desc        string    `db:"desc"`
	Image       string    `db:"image"`
	Genre       string    `db:"genre"`
	ReleaseDate string    `db:"release_date"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

func (g Game) GetAll() ([]*Game, error) {
	db := db.GetDB()
	sql := "SELECT * FROM games"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	log.Println(rows)

	var games []*Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.ID, &game.Name, &game.Desc, &game.Image, &game.Genre, &game.ReleaseDate, &game.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		games = append(games, &game)
	}
	return games, nil
}

func (g Game) GetByID(id string) (*Game, error) {
	db := db.GetDB()
	sql := "SELECT * FROM games WHERE id = " + id

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	log.Println(rows)
	game := Game{}
	for rows.Next() {
		if err := rows.Scan(&game.ID, &game.Name, &game.Desc, &game.Image, &game.Genre, &game.ReleaseDate, &game.UpdatedAt); err != nil {
			log.Fatal(err)
		}
	}
	return &game, nil
}

func (g Game) Create(game *Game) (*Game, error) {
	db := db.GetDB()
	id := uuid.NewV4()
	game.ID = id
	game.UpdatedAt = time.Now()
	stmt := `INSERT INTO games (id, name, desc, image, genre, releaseDate, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(stmt, game.ID, game.Name, game.Desc, game.Image, game.Genre, game.ReleaseDate, game.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when try to save data to database")
	}
	return game, nil
}

func (g Game) Delete(id string) error {
	db := db.GetDB()
	sql := "DELETE FROM games WHERE id = ?"
	_, err := db.Exec(sql, id)
	if err != nil {
		log.Println(err)
		return errors.New("error when try to delete " + id + " data from database")
	}
	return nil
}

func (h Game) Create2(payload form.GameCreate) (*Game, error) {
	db := db.GetDB()
	sql := `INSERT INTO games (name, desc, image, genre, releaseDate, updatedAt) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(sql, payload.Name, payload.Desc, payload.Image, payload.Genre, payload.ReleaseDate, strconv.FormatInt(time.Now().UnixNano(), 10))
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when try to save data to database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when try to retrieve last insert id")
	}

	// Query the database for the newly inserted game
	row := db.QueryRow("SELECT * FROM games WHERE id = ?", id)
	newGame := &Game{}
	err = row.Scan(&newGame.ID, &newGame.Name, &newGame.Desc, &newGame.Image, &newGame.Genre, &newGame.ReleaseDate, &newGame.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when try to retrieve new game from database")
	}

	return newGame, nil
}
