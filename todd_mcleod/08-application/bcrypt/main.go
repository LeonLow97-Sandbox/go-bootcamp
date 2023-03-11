package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	loginPwd1 := `password1234`
	
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd1))
	if err != nil {
		fmt.Println("You are not authorized!")
		return
	}
	fmt.Println("You are logged in!")
}