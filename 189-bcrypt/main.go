package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(hashedPass)

	loginPass := `password1234`

	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(loginPass))

	if err != nil {
		fmt.Println("You cannot Login")
		return
	}
	fmt.Println("You are logged in!")
}
