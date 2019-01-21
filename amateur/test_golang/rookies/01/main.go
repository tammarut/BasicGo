package main

//! Write Go_test func sumint(easy)
import "fmt"

func main() {
	fmt.Println("1+2=", sumint(1, 2))
	fmt.Println("1+5=", sumint(1, 5))
	fmt.Println("1+9=", sumint(1, 9))
}

func sumint(xi ...int) (sum int) {
	for _, v := range xi {
		sum += v
	}
	return sum
}
