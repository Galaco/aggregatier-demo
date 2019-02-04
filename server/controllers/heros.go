package controllers

import (
	"github.com/galaco/aggregatier/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AllHeroes(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	gameId,err := strconv.ParseInt(c.Param("gameId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H {
			"message": err,
		})
		return
	}

	heroes,err := models.AllHeroes(int(gameId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H {
			"message": err,
		})
		return
	}

	heroesJson := []gin.H{}
	for _,hero := range heroes {
		heroesJson = append(heroesJson, gin.H {
			"id" : hero.Id,
			"name" : hero.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H {
		"message":heroesJson,
	})
}
