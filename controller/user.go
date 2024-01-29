package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mistadave/gt-app/form"
	"github.com/mistadave/gt-app/models"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) GetUsers(c *gin.Context) {
	users, err := userModel.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (u UserController) GetById(c *gin.Context) {
	log.Println("Get user by id")
	id := c.Params.ByName("id")
	user, err := userModel.GetByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (u UserController) UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user form.UserUpdate
	c.BindJSON(&user)
	userUpdated, err := userModel.Update(id, user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userUpdated)
	}
}

func (u UserController) CreateUser(c *gin.Context) {
	var signUp form.UserSignup
	c.BindJSON(&signUp)
	user, err := userModel.Signup(signUp)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
