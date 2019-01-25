package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AllLists(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H {
		"message":"Not implemented yet, will return all lists",
	})
}

func GetList(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H {
		"message":"Not implemented yet, will return a list",
	})
}


func NewList(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H {
		"message":"Not implemented yet, will write a list to db",
	})
}

func UpdateList(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H {
		"message":"Not implemented yet. Will update a list in db",
	})
}
