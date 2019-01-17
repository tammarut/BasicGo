package main

import (
	"errors"
	"log"
	"math"
)

//! Learn about error.New (create own error)
func main() {
	_, err := sqrt(-4)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, errors.New("Oh, can't get negative number.")
	}
	return math.Sqrt(input), nil
}
