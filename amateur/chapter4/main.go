package main

//! เขียนโปรแกรมinputเลข 2 ตัวแล้วแสดงค่าที่มากกว่าทางจอ
import (
	"fmt"
)

func max(first, second int) int {
	if first < second {
		return second
	} else if first > second {
		return first
	}
	return 0
}
func main() {
	var first, second int

	fmt.Print("Input first number: ")
	fmt.Scanln(&first)
	fmt.Print("Input second number: ")
	fmt.Scanln(&second)

	if result := max(first, second); result == 0 {
		fmt.Println("It's equal.")
	} else {
		fmt.Println("Max: ", result)
	}

}
func init() {
	fmt.Println("Hi, chapter4")
	fmt.Println("-------------------------------")
}
