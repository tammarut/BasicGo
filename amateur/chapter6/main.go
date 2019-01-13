package main

//! เขียนโปรแกรมinputเลข 1ตัว แล้วprintตั้งแต่1ถึงเลขที่ใส่
import "fmt"

func printFor(input *int) {
	for i := 1; i <= *input; i++ {
		fmt.Printf("%d ", i)
	}
}
func main() {
	var input int

	fmt.Print("Input destinated number: ")
	fmt.Scanln(&input)
	printFor(&input)
}
func init() {
	fmt.Println("Hi, chapter6")
	fmt.Println("----------------------")
}
