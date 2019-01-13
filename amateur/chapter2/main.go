package main

import "fmt"

//! Operation(+ - * / mod) =>first, second
func main() {
	var first, second int

	fmt.Print("Input number1: ")
	fmt.Scanln(&first)
	fmt.Print("Input number2: ")
	fmt.Scanln(&second)

	fmt.Println("Result of")
	fmt.Printf("%d + %d = %d \n", first, second, first+second)
	fmt.Printf("%d - %d = %d \n", first, second, first-second)
	fmt.Printf("%d * %d = %d \n", first, second, first*second)
	fmt.Printf("%d / %d = %d \n", first, second, first/second)
	fmt.Printf("%d mod %d = %d \n", first, second, first%second)
}
func init() {
	fmt.Println("Hi, Chapter2 ")
	fmt.Println("-------------------------------")
}
