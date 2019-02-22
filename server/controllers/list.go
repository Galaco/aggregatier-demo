package controllers

import (
	"github.com/galaco/aggregatier/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Tiers(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	gameId, err := strconv.ParseInt(c.Param("gameId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}

	tiers, err := models.AllTiers(int(gameId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}

	tiersJson := []gin.H{}
	for _, tier := range tiers {
		tiersJson = append(tiersJson, gin.H{
			"id":   tier.Id,
			"name": tier.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": tiersJson,
	})
}

func AllLists(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet, will return all lists",
	})
}

func GetList(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet, will return a list",
	})
}

func NewList(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet, will write a list to db",
	})
}

func UpdateList(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H{
		"message": "Not implemented yet. Will update a list in db",
	})
}
