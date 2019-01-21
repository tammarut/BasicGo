package main

//! Nomal func sum(variadic)
import "fmt"

func main() {
	fmt.Println("5+10=", sumint(5, 10))
}

func sumint(xi ...int) (sum int) {
	for _, v := range xi {
		sum += v
	}
	return sum
}
