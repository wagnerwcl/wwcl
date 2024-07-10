package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createAt  time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (user User) OutputUserDetails() { // Method
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *User) CleanUserName() { // Method
	user.firstName = ""
	user.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "---",
			createAt:  time.Now(),
		},
	}
}

func New(firstName, lastName, birthDate string) (*User, error) { // Contructor function
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth data need to be set")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createAt:  time.Now(),
	}, nil
}
