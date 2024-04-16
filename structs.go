package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	adminEmail := getUserData("Please enter your email: ")
	adminPassword := getUserData("Please enter your password: ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	adminUser, err := user.NewAdmin(adminEmail, adminPassword)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.OutputUserDetails()
	appUser.ClearUserNames()
	appUser.OutputUserDetails()

	adminUser.OutputUserDetails()
	adminUser.ClearUserNames()
	adminUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
