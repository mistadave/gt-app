package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mistadave/gt-app/internal/models"
)

type GameController struct{}

var gameModel = new(models.Game)

func (u GameController) GetGames(c *gin.Context) {
	games, err := gameModel.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, games)
	}
}

func (u GameController) GetGameById(c *gin.Context) {
	id := c.Params.ByName("id")
	game, err := gameModel.GetByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, game)
	}
}

func (u GameController) CreateGame(c *gin.Context) {
	// TODO validate input and maybe create a new model for the input
	var game models.Game
	c.BindJSON(&game)
	gameNew, err := gameModel.Create(&game)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gameNew)
	}
}
