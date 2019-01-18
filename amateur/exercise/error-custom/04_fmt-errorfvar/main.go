package main

import (
	"fmt"
	"log"
	"math"
)

//! set error as variable(Errmath) and print
func main() {
	_, err := sqrt(-4)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(input float64) (float64, error) {
	if input < 0 {
		ErrMath := fmt.Errorf("Oh, you put negative number: %v", input)
		return 0, ErrMath
	}
	return math.Sqrt(input), nil
}
