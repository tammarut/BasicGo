package main

//! เขียนโปรแกรมinput 10ตัว แล้วแล้วแสดงผมรวมทั้งหมดและแสดงค่าเฉลี่ย
import "fmt"

func main() {
	var input, sum float64

	for i := 1; i <= 10; i++ {
		fmt.Printf("Input number %d: ", i)
		fmt.Scanln(&input)
		sum += input
	}
	fmt.Println("Sum: ", sum)
	fmt.Println("Average: ", sum/10)
}
func init() {
	fmt.Println("Hi, chapter7")
	fmt.Println("--------------------")
}
