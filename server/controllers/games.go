package controllers

import (
	"github.com/galaco/aggregatier/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindGame(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	gameId, err := strconv.ParseInt(c.Param("gameId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}

	game, err := models.FindGame(int(gameId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": gin.H{
			"id":        game.Id,
			"name":      game.Name,
			"shortName": game.ShortName,
		},
	})
}

func AllGames(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	games, err := models.AllGames()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}

	gamesJson := []gin.H{}
	for _, game := range games {
		gamesJson = append(gamesJson, gin.H{
			"id":        game.Id,
			"name":      game.Name,
			"shortName": game.ShortName,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": gamesJson,
	})

}
