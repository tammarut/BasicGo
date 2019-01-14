package main

import "fmt"

//* Update! :removed .git folder
//! My Profile input from keyboard and print out
func main() {
	var firstname, lastname, gender string
	var age int
	var grade float32

	fmt.Print("firstname: ")
	fmt.Scanln(&firstname)

	fmt.Print("lastname: ")
	fmt.Scanln(&lastname)

	fmt.Print("gender(M/F): ")
	fmt.Scanln(&gender)

	fmt.Print("age: ")
	fmt.Scanln(&age)

	fmt.Print("grade: ")
	fmt.Scanln(&grade)

	fmt.Println("\nName: ", firstname, lastname)
	fmt.Println("Gender: ", gender)
	fmt.Println("Age: ", age)
	fmt.Println("Grade: ", grade)
}
func init() {
	fmt.Println("Hi,chapter 1")
	fmt.Println("---------------------------------")
}
