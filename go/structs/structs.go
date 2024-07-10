package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("admin@example.com", "test123")

	admin.OutputUserDetails()
	admin.CleanUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.CleanUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
