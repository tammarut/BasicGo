package main

import "fmt"

//! เขียนโปรแกรมinputเลข 1ตัว แล้วแสดงดอกจัน(*)เป็นสี่เหลี่ยม nxn แบบตารางหมากรุก
func main() {
	var keyboard int
	fmt.Print("Input number: ")
	fmt.Scanln(&keyboard)

	for i := 1; i <= keyboard; i++ {
		for j := 1; j <= keyboard; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
}
