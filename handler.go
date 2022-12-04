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

func buildProfile(login, userID string) string {
	uri := getCredentials()
	if registerProfile(uri, login, userID) != true {
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

func getProfilehandler(userID string) string {
	uri := getCredentials()
	profileFound := getUserProfile(uri, userID)

	if profileFound != nil {
		dest := Profile{}
	}

	return "f"
}

func searchUserhandler(username string) string {
	uri := getCredentials()
	results := searchUser(uri, username)

	fmt.Println(len(results))
	for _, result := range results {
		fmt.Println(result)
		fmt.Println("ID")
		fmt.Println(result["Id"])
	}
	return "nil"
}
