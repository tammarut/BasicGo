package main

//! เขียนโปรแกรมรับinputเลขจำนวนเต็มไปเรื่อยๆจนกว่าจะป้อน0 เมื่อรับ0แล้วโปรแกรมจะprintจำนวนตัวเลข
//! พร้อมผลรวมแยกเป็นเลขคู่(even)และเลขคี่(odd)
import "fmt"

func main() {
	var input, evencount, oddcount int
	evensum, oddsum := 0, 0
	for {
		fmt.Print("Input number: ")
		fmt.Scanln(&input)
		if input%2 == 0 {
			evensum += input
			evencount++
		} else {
			oddsum += input
			oddcount++
		}
		if input == 0 {
			fmt.Printf("Sum %d even numbers: %d\n", evencount-1, evensum)
			fmt.Printf("Sum %d odd numbers: %d\n", oddcount, oddsum)
			break
		}
	}
}
func init() {
	fmt.Println("Hi, chapter8")
	fmt.Println("----------------------------")
}
