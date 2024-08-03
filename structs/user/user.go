package user


import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}


func NewAdmin(email, password string) Admin {
	return Admin{
		email,
		password,
		User{
			"Admin",
			"Admin",
			"----",
			time.Now(),
		},
	}
}

func New(userFirstName, userLastName, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("First Name, Last Name and BirthDate are Required")
	}
	return &User{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}, nil
}


func (appUser *User) OutputUserDetails() {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthdate)
}

func (appUser *User) ClearUserName() {
	appUser.firstName = ""
	appUser.lastName = ""
}
