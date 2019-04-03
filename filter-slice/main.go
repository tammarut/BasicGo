package main

//! Algorithm filter array
import (
	"fmt"
)

func main() {
	hold1 := []int{}
	hold2 := []int{}
	config := []int{100, 20}
	actual := []int{1, 2, 20, 100}

	for _, c := range config {
		for _, a := range actual {
			if a == c {
				hold1 = append(hold1, a)
			} else {
				hold2 = append(hold2, a)
			}
		}
		actual = hold2
		hold2 = []int{}
	}

	result := append(hold1, actual...)
	fmt.Println(result)
}
