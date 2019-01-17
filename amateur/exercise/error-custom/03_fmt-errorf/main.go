package main

import (
	"fmt"
	"log"
)

//! Errorf =>เอาค่าออกมาให้เห็นลย
func main() {
	_, err := sqrt(-4)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, fmt.Errorf("norgate math again: negative number: %v", input)
	}
	return 42, nil
}
