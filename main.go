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

	r.GET("/login/:login/:psw", login)

	r.POST("/register/:login/:psw", register)

	r.Run()
}

func login(c *gin.Context) {
	login := c.Param("login")
	psw := c.Param("psw")

	if LoginHandler(login, psw) != true {
		c.JSON(200, gin.H{
			"failed": "404",
		})
	} else {
		c.JSON(200, gin.H{
			"success": "200",
		})
	}
}

func register(c *gin.Context) {
	login := c.Param("login")
	psw := c.Param("psw")

	resp := RegisterHandler(login, psw)

	if resp != "200" {
		c.JSON(200, gin.H{
			"failed": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"success": "200",
		})
	}
}
