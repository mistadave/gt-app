package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mistadave/gt-app/internal/models"
)

type TagController struct{}

var tagModel = new(models.Tag)

func (t TagController) GetTags(c *gin.Context) {
	log.Println("Get all links")
	links, err := tagModel.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, links)
	}
}

func (t TagController) SearchTag(c *gin.Context) {
	name := c.Params.ByName("name")
	links, err := tagModel.SearchBy(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, links)
	}
}

func (t TagController) CreateTag(c *gin.Context) {
	var tag models.Tag
	c.BindJSON(&tag)
	tagNew, err := tagModel.Create(&tag)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, tagNew)
	}
}
