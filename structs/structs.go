package main

import (
	"fmt"
	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, error := user.New(userFirstName, userLastName, userBirthdate)

	admin:= user.NewAdmin("text@text.com","text123");
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	if error != nil {
		panic("Validation failed")
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
