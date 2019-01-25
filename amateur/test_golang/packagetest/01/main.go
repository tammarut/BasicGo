package main

//! Can call func sumint by reference folder
//! go test หรือ go test ./... =>Run test beneath current folder.
import (
	"fmt"

	"github.com/tammarut/amateur/test_golang/packagetest/acdc"
)

func main() {
	fmt.Println(acdc.Sumint(1, 2))
	fmt.Println(acdc.Sumint(1, 2, 10))
}
