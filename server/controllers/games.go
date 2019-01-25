package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllGames(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H {
		"message":"Not implemented. Will return all games",
	})
}
