//! Same rookie(easy) -> normal func sum
package main

import "fmt"

func main() {
	fmt.Println("1 + 2 =", sumint(1, 2))
}

func sumint(xi ...int) (sum int) {
	for _, v := range xi {
		sum += v
	}
	return sum
}
