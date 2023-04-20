package main

type Profile struct {
	FullName     string
	Description  string
	PhoneNb      string
	Email        string
	Notification []Notification
}

// Capital letters function to export them
func NewProfile(fullName, description, phoneNB, email string) Profile {
	product := Profile{fullName, description, phoneNB, email, []Notification{}}
	return product
}
