package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mistadave/gt-app/internal/models"
	uuid "github.com/satori/go.uuid"
)

type GameController struct{}

var gameModel = new(models.Game)
var genreModel = new(models.Genre)

func (u GameController) GetGenres(c *gin.Context) {
	genres, err := genreModel.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, genres)
	}
}

func (u GameController) CreateGenre(c *gin.Context) {
	var genre models.Genre
	c.BindJSON(&genre)
	genreNew, err := genreModel.Create(&genre)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, genreNew)
	}
}

func (u GameController) DeleteGenre(c *gin.Context) {
	id := c.Params.ByName("id")
	err := genreModel.Delete(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
	}
}

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

func (u GameController) UpdateGame(c *gin.Context) {
	var game models.Game
	id := c.Params.ByName("id")
	uuid, err := uuid.FromString(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&game)
	game.ID = uuid
	gameNew, err := gameModel.Update(&game)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gameNew)
	}
}
