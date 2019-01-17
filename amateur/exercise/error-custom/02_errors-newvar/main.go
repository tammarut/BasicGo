package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

//! kept error.New =>global variable
//* global variable easy to call
var ErrMath = errors.New("Oh, can't put negative number.")

func main() {
	fmt.Printf("%T \n", ErrMath)
	fmt.Println("----------------------")

	_, err := sqrt(-4)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, ErrMath
	}
	return math.Sqrt(input), nil
}
