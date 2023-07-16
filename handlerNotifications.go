package main

import (
	"fmt"
	"strings"
)

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
		Status:  Status,
	}

	err := AddNotification(uri, UserID, Notification)

	if !err {
		return "Failed to insert Notification"
	}
	return "Success"
}

func delNotificationHandler(UserID, Title string) string {
	uri := getCredentials()

	fmt.Println(Title)

	DelNotification(uri, UserID, Title)
	return "Success"
}
