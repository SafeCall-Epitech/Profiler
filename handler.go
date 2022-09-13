package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

func RegisterHandler(id, psw string) string {
	uri := getCredentials()
	users := GetUsers(uri)

	if len(id) < 5 {
		return "id too short" // Id too short
	}

	if len(psw) < 7 {
		return "password too short" // password too short
	}

	for _, info := range users {
		if info["login"] == id {
			return "Id already taken" // Id already taken
		}
	}

	if AddUser(uri, id, psw) != true {
		return "Unknown error"
	}
	return "200"
}

func LoginHandler(id, psw string) bool {
	uri := getCredentials()
	users := GetUsers(uri)

	for _, info := range users {
		if info["login"] == id && info["psw"] == psw {
			return true
		}
	}

	return false
}
