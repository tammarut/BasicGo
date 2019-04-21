package main

import "fmt"

func main() {
	A := []int{1, 2, 3}
	B := []int{4, 5, 6}
	sumArray := make([]int, len(A))

	for i := 0; i < 3; i++ {
		sumArray[i] = A[i] + B[i]
	}

	fmt.Println(sumArray)
}
