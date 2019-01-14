package main

//! เขียนโปรแกรมรับเลขบวกinput 1ตัว แล้วแสดงsquare root หากใส่ติดลบให้ป้อนใหม่จนกว่าจะใส่บวก
import (
	"fmt"
	"math"
)

func main() {
	var input int
	for {
		fmt.Print("Input non-negative number: ")
		fmt.Scanln(&input)
		if input >= 0 {
			break
		}
	}

	fmt.Printf("Square root %d = %.2f \n", input, math.Sqrt(float64(input)))
}
