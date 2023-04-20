package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
		addDelFriend(uri, userID, "?"+dest, "$pull")
	} else if action == "accept" {
		addDelFriend(uri, userID, "?"+dest, "$pull")
		addDelFriend(uri, userID, dest, "$push")
	} else if action == "deny" {
		addDelFriend(uri, userID, "?"+dest, "$pull")
		addDelFriend(uri, dest, userID, "$pull")
	}

	return "200"
}

func getFriendsHandler(userID string) []string {
	uri := getCredentials()
	result := GetFriends(uri, userID)

	str := fmt.Sprintf("%v", result)
	dest := strings.Split(str[1:len(str)-1], " ")
	return dest
}

func deleteUserData(userID string) string {
	uri := getCredentials()
	_ = deleteUserProfile(uri, userID)
	return "success"
}

func addEventHandler(guest1, guest2, subject, date string) string {
	uri := getCredentials()
	event := Event{
		Guests:    guest1 + "&" + guest2,
		Subject:   subject,
		Date:      date,
		Confirmed: false,
	}
	err := AddEvent(uri, guest1, event)
	if !err {
		return "Failed to insert event"
	}
	err = AddEvent(uri, guest2, event)

	if !err {
		return "Failed to insert event"
	}
	return "Success"
}

func delEventHandler(guest1, guest2, date string) string {
	uri := getCredentials()

	fmt.Println(guest1, date)

	DelEvent(uri, guest1, date)
	DelEvent(uri, guest2, date)
	return "Success"
}

func listEventHandler(userID string) []Event {
	uri := getCredentials()
	profileFound := getUserProfile(uri, userID)
	var events []Event

	test := profileFound["Agenda"]
	a := test.(primitive.A)

	for _, v := range a {
		b := v.(primitive.M)
		bi, _ := strconv.ParseBool(fmt.Sprint(b["Confirmed"]))

		event := Event{
			Guests:    fmt.Sprint(b["Guests"]),
			Subject:   fmt.Sprint(b["Subject"]),
			Date:      fmt.Sprint(b["Date"]),
			Confirmed: bi,
		}
		events = append(events, event)
	}
	return events
}

//Notification handling

// get notification of specified user
func GetNotification(userID string) []string {
	uri := getCredentials()
	result := GetNotificationsProfile(uri, userID)

	str := fmt.Sprintf("%v", result)
	dest := strings.Split(str[1:len(str)-1], " ")
	return dest[1:]
}

func addNotificationHandler(UserID, Title, Content string, Status bool) string {
	uri := getCredentials()
	Notification := Notification{
		Title:   Title,
		Content: Content,
		Status:  false,
	}
	err := AddNotification(uri, UserID, Notification)
	if !err {
		return "Failed to insert event"
	}
	err = AddNotification(uri, UserID, Notification)

	if !err {
		return "Failed to insert Notification"
	}
	return "Success"
}

func delNotificationHandler(UserID, Title, Content string) string {
	uri := getCredentials()

	fmt.Println(Title, Content)

	DelNotification(uri, UserID, Title)
	DelNotification(uri, UserID, Title)
	return "Success"
}
