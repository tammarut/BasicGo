package main

//! เขียนโปรแกรมรับinputเลข 1ตัว แสดงผลว่าเป็นจำนวนเฉพาะ(Prime number)หรือไม่
import "fmt"

func main() {
	var input int
	fmt.Print("Input number: ")
	fmt.Scanln(&input)

	if input%2 != 0 && input%3 != 0 && input%5 != 0 || input == 2 || input == 3 || input == 5 {
		fmt.Println("It's a prime number.")
	} else {
		fmt.Println("It is not prime number.")
	}
}
