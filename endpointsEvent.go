package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddEventStruct struct {
	Guest1  string `bson:"Guest1"`
	Guest2  string `bson:"Guest2"`
	Subject string `bson:"Subject"`
	Date    string `bson:"Date"`
}

type DelEventStruct struct {
	Guest1 string `bson:"Guest1"`
	Guest2 string `bson:"Guest2"`
	Date   string `bson:"Date"`
}

// FIXME Add a check to see if the guests exist
func addEventEndpoint(c *gin.Context) {
	var data AddEventStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := addEventHandler(data.Guest1, data.Guest2, data.Subject, data.Date)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func delEventEndpoint(c *gin.Context) {
	var data DelEventStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := delEventHandler(data.Guest1, data.Guest2, data.Date)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func listEventEndpoint(c *gin.Context) {
	user := c.Param("userID")

	a := listEventHandler(user)

	c.JSON(200, gin.H{
		"Success ": a,
	})
}
