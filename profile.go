package main

type Profile struct {
	FullName    string
	Description string
	PhoneNb     string
	Email       string
}

// Capital letters function to export them
func NewProfile(fullName, description, phoneNB, email string) Profile {
	product := Profile{fullName, description, phoneNB, email}
	return product
}
