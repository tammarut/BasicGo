package main

//! เขียนโปรแกรมinputจนกว่าผลรวมที่ป้อน=100
import "fmt"

func main() {
	var input, sum int
	for {
		fmt.Print("Input number (until sum=100): ")
		fmt.Scanln(&input)
		sum = sum + input
		fmt.Println("Sum: ", sum)
		if sum == 100 {
			break
		}
	}

}
func init() {
	fmt.Println("Hi, chapter9")
	fmt.Println("-----------------------")
}
