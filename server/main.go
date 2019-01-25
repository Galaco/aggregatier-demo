package main

import (
	"github.com/galaco/aggregatier/controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/games/all", controllers.AllGames)

		api.GET("/heroes/all", controllers.AllHeroes)

		api.GET("/list/all", controllers.AllLists)
		api.GET("/list/get", controllers.GetList)
		api.POST("/list/create", controllers.NewList)
		api.POST("/list/edit", controllers.UpdateList)
	}

	// Start and run the server
	router.Run(":3000")
}