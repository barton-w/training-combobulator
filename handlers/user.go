package handlers

import (
	"fmt"
	"training-combobulator/models"
)

type dao interface {
	FindUserByName(first, last string) *models.User
}

func SomeHandler(u dao) {
	fmt.Println("Looking for the answer...")

	user := u.FindUserByName("Abby", "Smith")

	fmt.Println("I found this user:")
	fmt.Printf("%+v\n", user)
}
