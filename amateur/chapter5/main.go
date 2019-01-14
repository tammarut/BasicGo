package main

//! เขียนโปรแกรมinputเลข 1ตัวแล้วบอกว่าเป็นคู่(even)/คี่(Odd)
import "fmt"

func main() {
	var input int

	fmt.Print("Input number: ")
	fmt.Scanln(&input)

	if input == 0 {
		fmt.Println("It's zero.")
	} else if input%2 == 0 {
		fmt.Printf("%d is even.", input)
	} else {
		fmt.Printf("%d is odd.", input)
	}

}
func init() {
	fmt.Println("Hi, chapter5")
	fmt.Println("-----------------------------")
}
