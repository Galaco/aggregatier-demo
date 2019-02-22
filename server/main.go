package main

import (
	"github.com/galaco/aggregatier/controllers"
	"github.com/galaco/aggregatier/models"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	initRoutes(router)
	initServices()

	// Start and run the server
	router.Run(":3000")
}

func initRoutes(router *gin.Engine) {
	// CORS
	router.Use(cors.Default())

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/games/all", controllers.AllGames)
		api.GET("/game/:gameId", controllers.FindGame)
		api.GET("/game/:gameId/tiers", controllers.Tiers)

		api.GET("/heroes/all/:gameId", controllers.AllHeroes)

		api.GET("/list/all", controllers.AllLists)
		api.GET("/list/get", controllers.GetList)
		api.POST("/list/create", controllers.NewList)
		api.POST("/list/edit", controllers.UpdateList)
	}
}

func initServices() {
	models.InitDB("root:password@tcp(database)/aggregatier")
}