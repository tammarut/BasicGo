package main

import (
	"fmt"
	"log"
	"math"
)

//! Error use type struct and method

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("Error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-4)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(input float64) (float64, error) {
	if input < 0 {
		nme := fmt.Errorf("Math redux: negative number: %v", input)
		return 0, norgateMathError{"50.2 N", "99.4 W", nme}
	}
	return math.Sqrt(input), nil
}
