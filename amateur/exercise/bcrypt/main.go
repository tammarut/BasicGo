package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	secret := `123456`
	byteslice, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Before:", secret)
	fmt.Println("After:", byteslice)

	user1 := `123456`

	err = bcrypt.CompareHashAndPassword(byteslice, []byte(user1))
	if err != nil {
		fmt.Println("Wrong password!.")
		return
	}
	fmt.Println("Login! success.")

}
func init() {
	fmt.Println("Hi!, bcrypt golang.")
	fmt.Println("------------------------------")
}
