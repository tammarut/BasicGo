package main

import "fmt"

//! เขียนรับเลขจากตีย์บอร์ดแล้ววาดรูป * สี่เหลี่ยมกลวง nxn
func main() {
	var input int
	fmt.Print("Input number of stars: ")
	fmt.Scanln(&input)

	for round := 1; round <= input; round++ {
		for i := 1; i <= input; i++ {
			if round == 1 || round == input || i == 1 || i == input {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
