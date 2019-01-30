package main

//! It's your dog.
//! Struct canine
//! fido ->>Call func Years, YearsTwo
import (
	"fmt"

	"github.com/tammarut/amateur/ninja_lv13/01/dog" //=> Auto import
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
