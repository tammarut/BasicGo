package main

import "fmt"

func fibonacciFactory() func() int {
	a, b := 0, 1
	return func() int {
		defer func() { a, b = b, a+b }()
		return a
	}
}

func main() {
	fibonacci := fibonacciFactory()
	for i := 0; i <= 10; i++ {
		fmt.Println(fibonacci())
	}
} //! Run init first before main
func init() {
	fmt.Println("Hi, fibonacci closure.")
	fmt.Println("------------------------------")
}
