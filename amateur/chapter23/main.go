package main

//! เขียนโปรแกรมรับจำนวนจากคีย์บอด แล้ววาดสามเหลี่่ยมเต็ม nxn

import "fmt"

//. Soulution 1
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

//. Solution 2
// func main() {
// 	var input, x, y int
// 	fmt.Print("Input numbers:")
// 	fmt.Scanln(&input)

// 	for y = input; y > 0; y-- {
// 		for x = -input; x <= input; x++ {
// 			if y <= (input+x) && y <= input-x {
// 				fmt.Print("*")
// 			} else {
// 				fmt.Print("-")
// 			}
// 		}
// 		fmt.Println() //=> Next new row
// 	}

// }
