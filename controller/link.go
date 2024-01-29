package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mistadave/gt-app/models"
)

type LinkController struct{}

var linkModel = new(models.Link)

func (l LinkController) GetLinks(c *gin.Context) {
	log.Println("Get all links")
	links, err := linkModel.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, links)
	}
}

func (l LinkController) GetLinksByGame(c *gin.Context) {
	id := c.Params.ByName("id")
	log.Println("Get all links by game id: " + id)
	links, err := linkModel.GetByGameID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, links)
	}
}

func (l LinkController) GetLinkById(c *gin.Context) {
	id := c.Params.ByName("id")
	link, err := linkModel.GetByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, link)
	}
}

func (l LinkController) CreateLink(c *gin.Context) {
	var link models.Link
	c.BindJSON(&link)
	linkNew, err := linkModel.Create(&link)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, linkNew)
	}
}
