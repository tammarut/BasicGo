package main

//! เขียนโปรแกรมรับinputเลข 10ตัวแล้วแสดงเลขที่อยู่ระหว่างเลขคู่
import "fmt"

func main() {
	box := [10]int{}

	for i := 0; i < len(box); i++ {
		fmt.Printf("Input number %d: ", i+1)
		fmt.Scanln(&box[i])
	}
	for i := 1; i < len(box)-1; i++ {
		if box[i-1]%2 == 0 && box[i+1]%2 == 0 {
			fmt.Print(" ", box[i])
		}
	}

}
func init() {
	fmt.Println("Hi, chapter 12")

	fmt.Println("-------------------------")
}
