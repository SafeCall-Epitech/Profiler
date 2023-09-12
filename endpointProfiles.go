package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProfileStruct struct {
	Login string `bson:"Login"`
	Email string `bson:"Email"`
}

type EditDescriptionStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

type EditFullNameStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

type EditPhoneNBStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

type EditEmailNBStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

type EditProfilePicStruct struct {
	UserID string `bson:"UserID"`
	Data   string `bson:"Data"`
}

type DeleteUserStruct struct {
	UserID string `bson:"UserID"`
}

func searchUserEndpoint(c *gin.Context) {
	userID := c.Param("username")
	resp := searchUserhandler(userID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}

func getProfileEndpoint(c *gin.Context) {
	userID := c.Param("userID")
	resp := getProfilehandler(userID)

	c.JSON(200, gin.H{
		"Profile ": resp,
	})
}

func editEmail(c *gin.Context) {
	var data EditEmailNBStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := handleProfileEdition("Email", data.UserID, data.Data)
	if resp != "true" {
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
	var data EditPhoneNBStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := handleProfileEdition("PhoneNB", data.UserID, data.Data)
	if resp != "true" {
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
	var data EditFullNameStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := handleProfileEdition("FullName", data.UserID, data.Data)
	if resp != "true" {
		c.JSON(503, gin.H{
			"Internal error ": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "200",
		})
	}
}

func createProfile(c *gin.Context) {
	var data CreateProfileStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := buildProfile(data.Login, data.Email)
	if resp != "Success" {
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
	var data EditDescriptionStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := handleProfileEdition("Description", data.UserID, data.Data)
	if resp != "true" {
		c.JSON(503, gin.H{
			"Internal error ": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "200",
		})
	}
}

func editProfilePic(c *gin.Context) {
	var data EditProfilePicStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := handleProfileEdition("ProfilePic", data.UserID, data.Data)
	if resp != "true" {
		c.JSON(503, gin.H{
			"Internal error ": resp,
		})
	} else {
		c.JSON(200, gin.H{
			"Success ": "200",
		})
	}
}

func deleteUser(c *gin.Context) {
	var data DeleteUserStruct
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := deleteUserData(data.UserID)

	c.JSON(200, gin.H{
		"Success ": resp,
	})
}
