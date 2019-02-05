package main

//! เขียนโปรแกรมรับจำนวนจากคีย์บอด แล้ววาดDiamond height = n
import "fmt"

func main() {
	var input int
	fmt.Print("Input numbers: ")
	fmt.Scanln(&input)

	//* top diamond
	for row := 0; row < input; row++ {
		for j := 0; j <= (input * 2); j++ {
			if j >= input-row && j <= input+row {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//* bottom diamond
	for row := input; row >= 0; row-- {
		for j := 0; j <= (input * 2); j++ {
			if j >= input-row && j <= input+row {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
