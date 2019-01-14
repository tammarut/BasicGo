package main

//! เขียนโปรแกรมรับเลขinput 1ตัวแล้วแสดงดอกจัน(*)ตามจำนวนinputที่ป้อน
import "fmt"

func main() {
	var input int
	fmt.Print("Input amount of stars: ")
	fmt.Scanln(&input)
	for i := 0; i < input; i++ {
		fmt.Print("*")
	}
}
