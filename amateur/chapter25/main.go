package main

//! เขียนโปรแกรมรับเลขจำนวนเต็ม แล้วoutputเป็นเลขอังกฤษ(one two three four..)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var keyboard int
	fmt.Print("Input number: ")
	fmt.Scanln(&keyboard)

	if typeof(keyboard) == "int" {
		slicestring := strings.Split(strconv.Itoa(keyboard), "") //.Here is key

		for _, v := range slicestring {
			var hold string
			switch v {
			case "0":
				hold = "zero"
			case "1":
				hold = "one"
			case "2":
				hold = "two"
			case "3":
				hold = "three"
			case "4":
				hold = "four"
			case "5":
				hold = "five"
			case "6":
				hold = "six"
			case "7":
				hold = "seven"
			case "8":
				hold = "eight"
			case "9":
				hold = "nine"
			default:
				fmt.Println("Don't know!")
			}
			fmt.Printf("%s ", hold)
		}
	} else {
		fmt.Println("Please key number!")
	}
}
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
