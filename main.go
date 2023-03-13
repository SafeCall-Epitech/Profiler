package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	zmq "github.com/pebbe/zmq4"
)

// This function is here for test purpose with Postman
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORS())

	r.GET("/Profile/:userID", getProfileEndpoint)
	r.GET("/search/:username", searchUserEndpoint)

	r.POST("/create/:login/:email", createProfile)
	// r.POST("/create/:login/:userID", createProfile)
	r.POST("/description/:userID/:description", editDescription)
	r.POST("/FullName/:userID/:FullName", editFullName)
	r.POST("/PhoneNB/:userID/:PhoneNB", editPhoneNB)
	r.POST("/Email/:userID/:email", editEmail)

	r.POST("/friend/:userID/:dest/:action", actionFriend)
	r.POST("/friendRequest/:userID/:dest/:action", ManageRequest)
	r.GET("/friends/:userID", GetFriendsEndpoint)

	r.GET("/testZMQServer", server)

	r.Run(":8081")
}

func server(c *gin.Context) {
	fmt.Println("ready sir")
	//  Socket to talk to clients
	responder, _ := zmq.NewSocket(zmq.PAIR)
	defer responder.Close()
	// responder.Bind("ipc://test1")
	responder.Bind("tcp://*:5555")

	for {
		//  Wait for next request from client
		msg, _ := responder.Recv(0)
		fmt.Println("Received ", msg)

		//  Do some 'work'
		time.Sleep(time.Second)

		//  Send reply back to client
		reply := "World"
		responder.Send(reply, 0)
		fmt.Println("Sent ", reply)
	}
}

func createProfile(c *gin.Context) {
	login := c.Param("login")
	email := c.Param("email")
	resp := buildProfile(login, email)
	if resp != "success" {
		c.JSON(200, gin.H{
			"Internal error ": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "200",
		})
	}
}

func editDescription(c *gin.Context) {
	userID := c.Param("userID")
	description := c.Param("description")

	resp := handleProfileEdition("Description", userID, description)
	if resp != "success" {
		c.JSON(503, gin.H{
			"Internal error ": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "200",
		})
	}
}

func editFullName(c *gin.Context) {
	userID := c.Param("userID")
	FullName := c.Param("FullName")

	resp := handleProfileEdition("FullName", userID, FullName)
	if resp != "success" {
		c.JSON(503, gin.H{
			"Internal error ": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "200",
		})
	}
}

func editPhoneNB(c *gin.Context) {
	userID := c.Param("userID")
	PhoneNB := c.Param("PhoneNB")

	resp := handleProfileEdition("PhoneNB", userID, PhoneNB)
	if resp != "success" {
		c.JSON(503, gin.H{
			"Internal error ": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "200",
		})
	}
}

func editEmail(c *gin.Context) {
	userID := c.Param("userID")
	Email := c.Param("email")

	resp := handleProfileEdition("Email", userID, Email)
	if resp != "success" {
		c.JSON(503, gin.H{
			"Internal error ": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "200",
		})
	}
}

func getProfileEndpoint(c *gin.Context) {
	userID := c.Param("userID")
	resp := getProfilehandler(userID)

	c.JSON(200, gin.H{
		"Profile ": resp,
	})
}

func searchUserEndpoint(c *gin.Context) {
	userID := c.Param("username")
	resp := searchUserhandler(userID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func actionFriend(c *gin.Context) {
	userID := c.Param("userID")
	dest := c.Param("dest")
	action := c.Param("action")

	resp := actionFriendHandler(userID, dest, action)

	if resp != "200" {
		c.JSON(503, gin.H{
			"Not found ": "404",
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "You " + action + " your friend",
		})
	}
}

func ManageRequest(c *gin.Context) {
	userID := c.Param("userID")
	dest := c.Param("dest")
	action := c.Param("action")

	actionFriendHandler(userID, dest, action)

	c.JSON(200, gin.H{
		"Success ": action + "ed",
	})
}

func GetFriendsEndpoint(c *gin.Context) {
	userID := c.Param("userID")
	friends := getFriendsHandler(userID)

	c.JSON(200, gin.H{
		"Success ": friends,
	})
}
