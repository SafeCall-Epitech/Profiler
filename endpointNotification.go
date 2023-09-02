package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddNotificationStruct struct {
	UserID  string `bson:"UserID"`
	Title   string `bson:"Title"`
	Content string `bson:"Content"`
	Status  string `bson:"Status"`
}

type DelNotificationStruct struct {
	UserID string `bson:"UserID"`
	Title  string `bson:"Title"`
}

func GetUserNotification(c *gin.Context) {
	userID := c.Param("UserID")
	resp := GetNotification(userID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func addNotificationEndpoint(c *gin.Context) {
	var data AddNotificationStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := strconv.ParseBool(data.Status)

	if err != nil {
		c.JSON(200, gin.H{
			"Success ": err,
		})
	}

	resp := addNotificationHandler(data.UserID, data.Title, data.Content, res)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func delNotificationEndpoint(c *gin.Context) {
	var data DelNotificationStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := delNotificationHandler(data.UserID, data.Title)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}
