package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllHeroes(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H {
		"message":"Not implemented yet. Will return all heroes for a specific game",
	})
}
