package main

import (
	"github.com/gin-gonic/gin"
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

	r.POST("/create", createProfile)
	r.POST("/Description", editDescription)
	r.POST("/FullName", editFullName)
	r.POST("/PhoneNB", editPhoneNB)
	r.POST("/Email", editEmail)
	r.POST("/ProfilePic", editProfilePic)
	r.POST("/delete", deleteUser)

	r.POST("/friend", actionFriend)
	r.POST("/friendRequest", ManageRequest)
	r.GET("/friends/:userID", GetFriendsEndpoint)

	r.POST("/addEvent", addEventEndpoint)
	r.POST("/delEvent", delEventEndpoint)
	r.GET("/listEvent/:userID", listEventEndpoint)

	r.POST("/AddNotification", addNotificationEndpoint)
	r.POST("/DelNotification", delNotificationEndpoint)
	r.GET("/notification/:UserID", GetUserNotification)

	r.GET("/testZMQServer", server)

	r.Run(":8081")
}

// KEPT FOR TEST PURPOSES
func server(c *gin.Context) {
	// fmt.Println("ready sir")
	// //  Socket to talk to clients
	// responder, _ := zmq.NewSocket(zmq.PAIR)
	// defer responder.Close()
	// // responder.Bind("ipc://test1")
	// responder.Bind("tcp://*:5555")

	// for {
	// 	//  Wait for next request from client
	// 	msg, _ := responder.Recv(0)
	// 	fmt.Println("Received ", msg)

	// 	//  Do some 'work'
	// 	time.Sleep(time.Second)

	// 	//  Send reply back to client
	// 	reply := "World"
	// 	responder.Send(reply, 0)
	// 	fmt.Println("Sent ", reply)
	// }
}
