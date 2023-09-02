package main

import (
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addEventHandler(guest1, guest2, subject, date string) string {
	uri := getCredentials()
	event := Event{
		Guests:    guest1 + "+" + guest2,
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

	if test != nil {
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
	return nil
}
