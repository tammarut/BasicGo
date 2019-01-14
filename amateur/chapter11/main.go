package main

//! เขียนโปรแกรมinputเลข 3ตัว แสดงผลว่าเป็นสามเหลี่ยมมุมฉากหรือไม่?
import (
	"fmt"
)

func main() {
	var x, y, z int

	fmt.Print("X: ")
	fmt.Scanln(&x)
	fmt.Print("Y: ")
	fmt.Scanln(&y)
	fmt.Print("Z: ")
	fmt.Scanln(&z)

	if (x*x)+(y*y) == (z*z) || (y*y)+(z*z) == (x*x) || (x*x)+(z*z) == (y*y) {
		fmt.Println("X, Y, Z are side of right triangle.")
	} else {
		fmt.Println("No, it's not side of right triangle. ")
	}
}
func init() {
	fmt.Println("Hi, chapter11: Side triangle")
	fmt.Println("--------------------------------")
}
