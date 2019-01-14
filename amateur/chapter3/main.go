package main

import (
	"fmt"
	"math"
)

//! Ax^2+Bx+C =0
func main() {
	var A, B, C float64

	fmt.Println("Ax^2 + Bx +C =0")
	fmt.Print("Input A: ")
	fmt.Scanln(&A)
	fmt.Print("Input B: ")
	fmt.Scanln(&B)
	fmt.Print("Input C: ")
	fmt.Scanln(&C)

	inroot := math.Pow(B, 2) - 4*(A*C)
	xup := (-B + math.Sqrt(inroot)) / (2 * A)
	xdown := (-B - math.Sqrt(inroot)) / (2 * A)

	fmt.Printf("\n%.2fx^2 + %.2fx + %.2f =0\n", A, B, C)
	fmt.Printf("Anwser1: %.2f \n", xup)
	fmt.Printf("Answer2: %.2f", xdown)

}
func init() {
	fmt.Println("Hi, chapter 3!")
	fmt.Println("-----------------------------------------")
}
