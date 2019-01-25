package main

//! Both test
import (
	"fmt"

	"github.com/tammarut/amateur/test_golang/packagetest/02/mathtwo"
)

func main() {
	fmt.Println("1+2 =", mathtwo.Sumint(1, 2))
	fmt.Println("2+5 =", mathtwo.Sumint(2, 5))
	fmt.Println("1+10 =", mathtwo.Sumint(1, 10))

}
