package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActionFriendStruct struct {
	UserID  string `bson:"UserID"`
	Dest    string `bson:"Dest"`
	Subject string `bson:"Subject"`
	Action  string `bson:"Action"`
}

type ManageStruct struct {
	UserID  string `bson:"UserID"`
	Dest    string `bson:"Dest"`
	Subject string `bson:"Subject"`
	Action  string `bson:"Action"`
}

func actionFriend(c *gin.Context) {
	var data ActionFriendStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := AddFriendHandler(data)

	if resp != "200" {
		c.JSON(503, gin.H{
			"Not found ": "404",
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "You " + data.Action + " your friend",
		})
	}
}

func ManageRequest(c *gin.Context) {
	var data ManageStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := acceptFriendHandler(data)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func GetFriendsEndpoint(c *gin.Context) {
	userID := c.Param("userID")
	friends := listFriendHandler(userID)
	c.JSON(200, gin.H{
		"Success ": friends,
	})
}
