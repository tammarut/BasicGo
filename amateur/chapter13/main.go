package main

//! เขียนโปรแกรมรับinputเลข 10ตัวแล้วแสดงผลเรียงลำดับจากน้อย-มาก
import (
	"fmt"
)

func main() {
	box := [10]int{}
	max := 0
	for i := 0; i < len(box); i++ {
		fmt.Printf("Input number %d: ", i)
		fmt.Scanln(&box[i])
	}
	for index := 9; index > 0; index-- {
		for i := 0; i < index; i++ {
			if box[i] > box[i+1] {
				max = box[i]
				box[i] = box[i+1]
				box[i+1] = max
			}
		}
	}
	fmt.Print("Sorted: ")
	for j := 0; j < len(box); j++ {
		fmt.Printf("%d ", box[j])
	}
}
func init() {
	fmt.Println("Hi, chapter13")
	fmt.Println("---------------------------")
}
