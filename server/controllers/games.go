package controllers

import (
	"github.com/galaco/aggregatier/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllGames(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	games,err := models.AllGames()
	if err != nil {
		c.JSON(http.StatusOK, gin.H {
			"message": err,
		})
		return
	}

	gamesJson := []gin.H{}
	for _,game := range games {
		gamesJson = append(gamesJson, gin.H {
			"id" : game.Id,
			"name" : game.Name,
			"shortName" : game.ShortName,
		})
	}

	c.JSON(http.StatusOK, gin.H {
		"message": gamesJson,
	})

}
