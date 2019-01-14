package main

//! เขียนโปรแกรมinputเลข 1ตัว แล้วแสดงดอกจัน(*)เป็นสี่เหลี่ยม nxn
import "fmt"

func main() {
	var input int
	fmt.Print("Input n: ")
	fmt.Scanln(&input)

	for column := 0; column < input; column++ {
		for row := 0; row < input; row++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
