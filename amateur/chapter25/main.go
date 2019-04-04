package main

//! เขียนโปรแกรมรับเลขจำนวนเต็ม แล้วoutputเป็นเลขอังกฤษ(one two three four..)

import (
	"fmt"
	"strings"
)

func main() {
	var keyboard string
	fmt.Print("Input number: ")
	fmt.Scanln(&keyboard)

	slicestring := strings.Split(keyboard, "")

	for _, v := range slicestring {
		switch v {
		case "0":
			fmt.Print("zero ")
		case "1":
			fmt.Print("one ")
		case "2":
			fmt.Print("two ")
		case "3":
			fmt.Print("three ")
		case "4":
			fmt.Print("four ")
		case "5":
			fmt.Print("five ")
		case "6":
			fmt.Print("six ")
		case "7":
			fmt.Print("seven ")
		case "8":
			fmt.Print("eight ")
		case "9":
			fmt.Print("nine ")
		default:
			fmt.Println("Don't know!")
		}
	}
}
