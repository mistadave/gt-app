package models

import (
	"errors"
	"log"
	"time"

	"github.com/mistadave/gt-api/db"
	"github.com/mistadave/gt-api/form"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Birthday  string    `db:"birthday" json:"birthday"`
	Password  string    `db:"password" json:"-"`
	Gender    string    `db:"gender" json:"gender"`
	PhotoURL  string    `db:"photo_url" json:"photoUrl"`
	Time      int       `db:"time" json:"time"`
	Active    bool      `db:"active" json:"active"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

func (h User) Signup(userPayload form.UserSignup) (*User, error) {
	db := db.GetDB()
	id := uuid.NewV4()
	user := User{
		ID:        id,
		Name:      userPayload.Name,
		Birthday:  userPayload.BirthDay,
		Gender:    userPayload.Gender,
		PhotoURL:  userPayload.PhotoURL,
		Time:      0,
		Active:    true,
		UpdatedAt: time.Now(),
	}
	stmt := `INSERT INTO users (
		id, name, username, email, birthday, password, gender, photo_url, time, active) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(stmt, user.ID, user.Name, user.Birthday, user.Gender, user.PhotoURL, user.Time, user.Active)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when try to save data to database")
	}
	return &user, nil
}

func (h User) Login(userPayload form.UserLogin) (*User, error) {
	db := db.GetDB()
	sql := "SELECT * FROM users WHERE username = ? AND password = ?"
	rows, err := db.Query(sql, userPayload.Username, userPayload.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	user := User{}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.UpdatedAt); err != nil {
			log.Fatal(err)
		}
	}
	return &user, nil
}

func (h User) GetByID(id string) (*User, error) {
	db := db.GetDB()
	log.Println(`get user by id`, id)
	UserUUID := uuid.FromStringOrNil(id)
	if UserUUID == uuid.Nil {
		log.Println(`invalid uuid`)
		return nil, errors.New("invalid uuid")
	}
	sql := "SELECT * FROM users WHERE id = ?"
	rows, err := db.Query(sql, UserUUID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	user := User{}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Birthday, &user.Password, &user.Gender, &user.PhotoURL, &user.Time, &user.Active, &user.UpdatedAt); err != nil {
			log.Fatal(err)
		}
	}
	return &user, nil
}

func (u User) Update(id string, userPayload form.UserUpdate) (*User, error) {
	db := db.GetDB()
	UserUUID := uuid.FromStringOrNil(id)
	if UserUUID == uuid.Nil {
		return nil, errors.New("invalid uuid")
	}
	sql := "UPDATE users SET name = ?, email = ?, birthday = ?, gender = ?, photo_url = ?, time = ? WHERE id = ?"
	_, err := db.Exec(sql, userPayload.Name, userPayload.Email, userPayload.BirthDay, userPayload.Gender, userPayload.PhotoURL, userPayload.Time, id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when try to update " + id + " data from database")
	}
	return &User{}, nil
}

func (h User) GetAll() ([]*User, error) {
	db := db.GetDB()
	stmt := "SELECT id, name, username, email, birthday, password, gender, photo_url, time, active, updated_at FROM users"
	rows, err := db.Query(stmt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Birthday, &user.Password, &user.Gender, &user.PhotoURL, &user.Time, &user.Active, &user.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}
	return users, nil
}
