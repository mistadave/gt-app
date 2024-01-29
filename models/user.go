package models

import (
	"errors"
	"log"

	"github.com/mistadave/gt-app/form"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:not null"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Username string    `gorm:"type:varchar(100);unique;not null"`
	Email    string    `gorm:"type:varchar(255);unique;not null"`
	Birthday string    `gorm:"type:date"`
	Password string    `gorm:"type:varchar(300);not null"`
	Gender   string    `gorm:"type:varchar(10)"`
	PhotoURL string    `gorm:"type:varchar(255)"`
	Time     int       `gorm:"type:int"`
	Active   bool      `gorm:"type:boolean;default:true"`
	gorm.Model
}

func (h User) Signup(userPayload form.UserSignup) (*User, error) {
	db := GetDB()
	id := uuid.NewV4()
	user := User{
		ID:       id,
		Name:     userPayload.Name,
		Birthday: userPayload.BirthDay,
		Gender:   userPayload.Gender,
		PhotoURL: userPayload.PhotoURL,
		Time:     0,
		Active:   true,
	}

	err := db.Create(&user).Error
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when trying to save data to database")
	}
	return &user, nil
}

func (h User) Login(userPayload form.UserLogin) (*User, error) {
	db := GetDB()
	var user User
	result := db.Where("username = ? AND password = ?", userPayload.Username, userPayload.Password).First(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return &user, nil
}

func (h User) GetByID(id string) (*User, error) {
	db := GetDB()
	log.Println(`get user by id`, id)
	UserUUID := uuid.FromStringOrNil(id)
	if UserUUID == uuid.Nil {
		log.Println(`invalid uuid`)
		return nil, errors.New("invalid uuid")
	}

	var user User
	result := db.Where("id = ?", UserUUID).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	} else if result.Error != nil {
		// some other problem
		log.Fatal(result.Error)
	}
	return &user, nil
}

func (u User) Update(id string, userPayload form.UserUpdate) (*User, error) {
	db := GetDB()
	UserUUID := uuid.FromStringOrNil(id)
	if UserUUID == uuid.Nil {
		return nil, errors.New("invalid uuid")
	}

	user := User{
		Name:     userPayload.Name,
		Email:    userPayload.Email,
		Birthday: userPayload.BirthDay,
		Gender:   userPayload.Gender,
		PhotoURL: userPayload.PhotoURL,
		Time:     userPayload.Time,
	}

	result := db.Model(&User{}).Where("id = ?", UserUUID).Updates(user)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, errors.New("error when try to update " + id + " data from database")
	}
	return &User{}, nil
}

func (h User) GetAll() ([]*User, error) {
	db := GetDB()
	var users []*User
	result := db.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return users, nil
}
