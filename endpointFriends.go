package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActionFriendStruct struct {
	userID string `bson:"userID"`
	Dest   string `bson:"dest"`
	Action string `bson:"action"`
}

type ManageStruct struct {
	userID string `bson:"userID"`
	Dest   string `bson:"dest"`
	Action string `bson:"action"`
}

func actionFriend(c *gin.Context) {
	var data ActionFriendStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := actionFriendHandler(data.userID, data.Dest, data.Action)

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

	actionFriendHandler(data.userID, data.Dest, data.Action)

	c.JSON(200, gin.H{
		"Success ": data.Action + "ed",
	})
}

func GetFriendsEndpoint(c *gin.Context) {
	userID := c.Param("userID")
	friends := getFriendsHandler(userID)

	c.JSON(200, gin.H{
		"Success ": friends,
	})
}
