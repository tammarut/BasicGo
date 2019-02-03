package main

import "fmt"

//! เขียนโปรแกรมรับจำนวนจากคีย์บอด แล้ววาดสามเหลี่่ยมเต็ม nxn
func main() {
	var input int
	fmt.Print("Input numbers:")
	fmt.Scanln(&input)

	for row := 1; row <= input; row++ {
		for space := 1; space <= input-row; space++ {
			fmt.Print("-")
		}
		for star := 1; star <= (row*2)-1; star++ {
			fmt.Print("*")
		}

		fmt.Println() //=> Next new row
	}

}
