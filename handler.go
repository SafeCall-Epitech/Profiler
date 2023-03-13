package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Credentials struct {
	Uri string `json:"uri"`
}

// This function will get the uri in the json file to id to the db
func getCredentials() string {
	fileContent, err := os.Open("config.json")

	if err != nil {
		log.Fatal(err)
		return ""
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	res := Credentials{}
	json.Unmarshal([]byte(byteResult), &res)

	return res.Uri
}

func userToProto(username, psw string) UserMessage {
	user := UserMessage{
		Id:       1,
		Username: username,
		Password: psw,
		Settings: &SettingsMessage{
			DoNotDisturb: true,
			Language:     "eng",
		},
	}
	return user
}

func buildProfile(login, email string) string {
	uri := getCredentials()
	if !registerProfile(uri, login, email) {
		return "Internal error"
	}
	return "Success"
}

func handleProfileEdition(endpoint, userID, data string) string {
	uri := getCredentials()
	if endpoint == "Description" && len(data) > 350 {
		return "Too long description"
	}
	if endpoint == "FullName" && len(data) > 30 {
		return "Too long Full Name"
	}
	if endpoint == "PhoneNB" && len(data) > 15 {
		return "Too long PhoneNB"
	}
	if endpoint == "Email" && len(data) > 50 {
		return "Too long Email"
	}

	parsedData := strings.ReplaceAll(data, "_", " ") // FIXME
	_ = publishProfileUpdates(uri, endpoint, userID, parsedData)
	return "success"
}

func getProfilehandler(userID string) Profile {
	uri := getCredentials()
	profileFound := getUserProfile(uri, userID)

	if profileFound != nil {

		dest := NewProfile(
			fmt.Sprint(profileFound["FullName"]),
			fmt.Sprint(profileFound["Description"]),
			fmt.Sprint(profileFound["PhoneNB"]),
			fmt.Sprint(profileFound["Email"]),
		)
		return dest
	}

	return Profile{}
}

func searchUserhandler(username string) map[int]string {
	uri := getCredentials()
	results := searchUser(uri, username)
	m := make(map[int]string)

	for nb, result := range results {
		m[nb] = fmt.Sprintf(result["Id"].(string))
		if nb > 4 {
			return m
		}
	}
	return m
}

func actionFriendHandler(userID, dest, action string) string {
	found := getProfilehandler(dest)
	uri := getCredentials()

	if found.FullName == "" { // TODO check si il est déjà ami
		return "Not found"
	} else if action == "add" {
		addDelFriend(uri, userID, dest, "$push")
		addDelFriend(uri, dest, "?"+userID, "$push")
	} else if action == "rm" {
		addDelFriend(uri, dest, userID, "$pull")
		addDelFriend(uri, userID, dest, "$pull")
	} else if action == "accept" {
		addDelFriend(uri, userID, "?"+dest, "$pull")
		addDelFriend(uri, userID, dest, "$push")
	} else if action == "deny" {
		addDelFriend(uri, userID, "?"+dest, "$pull")
		addDelFriend(uri, dest, userID, "$pull")
	}

	return ""
}

func getFriendsHandler(userID string) []string {
	uri := getCredentials()
	result := GetFriends(uri, userID)

	str := fmt.Sprintf("%v", result)
	dest := strings.Split(str[1:len(str)-1], " ")
	return dest[1:]
}
