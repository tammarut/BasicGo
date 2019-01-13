package main

//! เขียนโปรแกรมinputเลข 3ตัวแล้วแสดงผลว่าเลขทั้งสามเป็นด้านของสามเหลี่ยมเหรือไม่
import "fmt"

func main() {
	var a, b, c, max int

	fmt.Print("A: ")
	fmt.Scanln(&a)
	fmt.Print("B: ")
	fmt.Scanln(&b)
	fmt.Print("C: ")
	fmt.Scanln(&c)

	if a > b {
		max = a
	}
	if c > max {
		max = c
	}
	if max+max < a+b+c {
		fmt.Println("A, B, C are side of triangle.")
	} else {
		fmt.Println("A, B, C are not side triangle.")
	}
}
func init() {
	fmt.Println("Hi, chapter10")
	fmt.Println("--------------------------")
}
