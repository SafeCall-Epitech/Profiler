package main

import (
	"fmt"
	"strings"
)

func actionFriendHandler(userID, dest, action string) string {
	found := getProfilehandler(dest)
	uri := getCredentials()

	if found.FullName == "" { // TODO check si il est déjà ami
		return "Not found"
	} else if action == "add" && !check_duplicata(userID, dest) {
		addDelFriend(uri, userID, dest, "$push")
		addDelFriend(uri, dest, "?"+userID, "$push")
	} else if action == "rm" {
		addDelFriend(uri, dest, userID, "$pull")
		addDelFriend(uri, userID, dest, "$pull")
		addDelFriend(uri, userID, "?"+dest, "$pull")
	} else if action == "accept" && !check_duplicata(userID, dest) {
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

func check_duplicata(userID, NewFriend string) bool {
	friends := getFriendsHandler(userID)
	for _, friend := range friends {
		if friend == NewFriend {
			return true
		}
	}
	return false
}
