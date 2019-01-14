package main

//! เขียนโปรแกรมรับinputเลข 1ตัวแล้วแสดงผลการแยกตัวประกอบ
import "fmt"

func main() {
	input, p := 0, 2
	fmt.Print("Input number: ")
	fmt.Scanln(&input)
	fmt.Printf("factor %d= ", input)

	for input > 1 && p < input {
		if input%p == 0 {
			fmt.Printf("%d x ", p)
			input = input / p
			p = 2
		} else {
			p++
		}
	}
	if input == 0 || input == 1 {
		fmt.Printf("%d ", input)
	} else {
		fmt.Printf("%d \n", p)
	}

}
