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

	r.POST("/create/:login/:userID", createProfile)
	r.POST("/description/:userID/:description", editDescription)
	r.POST("/FullName/:userID/:FullName", editFullName)

	r.Run(":8081")
}

func createProfile(c *gin.Context) {
	login := c.Param("login")
	id := c.Param("userID")
	resp := buildProfile(login, id)
	if resp != "success" {
		c.JSON(200, gin.H{
			"Internal error ": "503",
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

	resp := handleDescription("Description", userID, description)
	if resp != "success" {
		c.JSON(503, gin.H{
			"Internal error ": "503",
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

	resp := handleDescription("FullName", userID, FullName)
	if resp != "success" {
		c.JSON(503, gin.H{
			"Internal error ": "503",
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "200",
		})
	}
}
