package main

import (
	"fmt"
	"strconv"
)

func buildProfile(login, email string) string {
	uri := getCredentials()
	if !registerProfile(uri, login, email) {
		return "Internal error"
	}
	return "Success"
}

func handleProfileEdition(endpoint, userID, data string) string {
	fmt.Println(endpoint, "DATA : ", data)
	uri := getCredentials()
	if endpoint == "Description" && len(data) > 350 {
		return "Too long description"
	}
	if endpoint == "FullName" && len(data) > 30 {
		return "Too long Full Name"
	}
	if endpoint == "PhoneNb" && len(data) > 15 {
		return "Too long PhoneNB"
	}
	if endpoint == "Email" && len(data) > 50 {
		return "Too long Email"
	}
	if endpoint == "ProfilePic" && len(data) > 150 {
		return "Too long link"
	}

	resp := publishProfileUpdates(uri, endpoint, userID, data)
	return strconv.FormatBool(resp)

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

func deleteUserData(userID string) string {
	uri := getCredentials()
	_ = deleteUserProfile(uri, userID)
	return "success"
}
