package main

type Profile struct {
	FullName     string
	Description  string
	PhoneNb      string
	Email        string
	Notification []Notification
	ProfilePic   string
}

// Capital letters function to export them
func NewProfile(fullName, description, phoneNB, email, pic string) Profile {
	product := Profile{fullName, description, phoneNB, email, []Notification{}, pic}
	return product
}
