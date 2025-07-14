package main

type User struct {
	name    string
	message string
}

func NewUser(name string, message string) *User {
	return &User{
		name:    name,
		message: message,
	}
}
