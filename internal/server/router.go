package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mistadave/gt-app/internal/controller"
	"github.com/mistadave/gt-app/internal/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())
	health := new(controller.HealthController)
	router.GET("/health", health.Status)
	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	user := new(controller.UserController)
	game := new(controller.GameController)
	link := new(controller.LinkController)
	tag := new(controller.TagController)

	v1 := router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			authGroup.POST("/login", user.Login)
			// authGroup.POST("/register", user.Register)
		}
		v1.Use(middlewares.AuthMiddleware())
		userGroup := v1.Group("users")
		{
			// TODO find out if it's better to define inside or outside the group
			userGroup.GET("/", user.GetUsers)
			userGroup.GET("/:id", user.GetById)
			userGroup.PUT("/:id", user.UpdateUser)
			userGroup.POST("/", user.CreateUser)

		}
		gameGroup := v1.Group("games")
		{
			gameGroup.GET("/", game.GetGames)
			gameGroup.GET("/:id", game.GetGameById)
			gameGroup.POST("/", game.CreateGame)
			gameGroup.GET("/:id/links", link.GetLinksByGame)
		}
		linkGroup := v1.Group("links")
		{
			linkGroup.GET("/", link.GetLinks)
			linkGroup.GET("/:id", link.GetLinkById)
			linkGroup.POST("/", link.CreateLink)
		}
		tagGroup := v1.Group("tags")
		{
			tagGroup.GET("/", tag.GetTags)
			tagGroup.GET("/:name", tag.SearchTag)
			tagGroup.POST("/", tag.CreateTag)
		}
	}
	return router
}
