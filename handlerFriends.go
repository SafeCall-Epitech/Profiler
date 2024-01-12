package main

import (
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddFriendHandler(data ActionFriendStruct) string {
	found := getProfilehandler(data.Dest)
	uri := getCredentials()

	person0 := Friends{
		Id:      data.Dest,
		Subject: data.Subject,
		Active:  true,
	}
	person1 := Friends{
		Id:      data.UserID,
		Subject: data.Subject,
		Active:  true,
	}

	if found.FullName == "" {
		return "Not found"
	} else if data.Action == "add" && !check_duplicata(data.UserID, data.Dest) {
		person0.Active = false
		person1.Active = false
		person1.Id = "?" + person1.Id
		addDelFriendAdd(uri, data.UserID, "$push", person0)
		addDelFriendAdd(uri, data.Dest, "$push", person1)
	} else if data.Action == "delete" {
		addDelFriendAdd(uri, data.UserID, "$pull", person0)
		addDelFriendAdd(uri, data.Dest, "$pull", person1)
	}

	return "200"
}

func acceptFriendHandler(data ManageStruct) string {
	uri := getCredentials()
	friends := getFriendsFromID(data.UserID)
	person0 := Friends{
		Id:      data.Dest,
		Subject: data.Subject,
		Active:  true,
	}
	person1 := Friends{
		Id:      data.UserID,
		Subject: data.Subject,
		Active:  true,
	}
	if isFriendshipActivable(data.Dest, friends) && data.Action == "accept" {
		addDelFriendAdd(uri, data.UserID, "$push", person0)
		addDelFriendAdd(uri, data.Dest, "$push", person1)
		person0.Active = false
		person1.Active = false
		person0.Id = "?" + data.Dest
		addDelFriendAdd(uri, data.UserID, "$pull", person0)
		addDelFriendAdd(uri, data.Dest, "$pull", person1)
		return "accepted"
	} else if data.Action == "reject" {
		person0.Active = false
		person1.Active = false
		person0.Id = "?" + data.Dest
		addDelFriendAdd(uri, data.UserID, "$pull", person0)
		addDelFriendAdd(uri, data.Dest, "$pull", person1)
		return "rejected"
	} else {
		return "error : DEST " + data.Dest
	}
}

func isFriendshipActivable(dest string, friends []Friends) bool {
	for _, friend := range friends {
		fmt.Println(friend)
		if friend.Id == "?"+dest && !friend.Active {
			return true
		}
	}
	return false
}

func listFriendHandler(userID string) []Friends {
	uri := getCredentials()
	profileFound := getUserProfile(uri, userID)
	var friends []Friends
	test := profileFound["Friends"]

	if test != nil {
		a := test.(primitive.A)
		for _, v := range a {
			b := v.(primitive.M)
			bi, _ := strconv.ParseBool(fmt.Sprint(b["Active"]))

			friendProfile := getProfilehandler(fmt.Sprint(b["Id"]))

			friend := Friends{
				Id:       fmt.Sprint(b["Id"]),
				FullName: &friendProfile.FullName,
				Subject:  fmt.Sprint(b["Subject"]),
				Active:   bi,
			}
			friends = append(friends, friend)
		}
		return friends
	}
	return nil
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
		fmt.Println(friend, "friend")
		if len(friend) >= 3 {
			if friend[3:] == NewFriend {
				return true
			} else if friend[3:] == "?"+NewFriend {
				return true
			}
		}
	}
	return false
}
