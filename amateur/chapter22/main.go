package main

//! เขียนโปรแกรมรับจำนวนจากคีย์บอด แล้ววาดสามเหลี่่ยมครึ่งหงาย nxn
import "fmt"

func main() {
	var input int
	fmt.Print("Input number: ")
	fmt.Scanln(&input)

	for column := 1; column <= input; column++ {
		for sp := input; sp > column; sp-- {
			fmt.Print("-")
		}
		for star := 1; star <= column; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
