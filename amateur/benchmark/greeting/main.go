package main

//! Call func Greet by reference folder saying
import (
	"fmt"

	"github.com/tammarut/amateur/benchmark/00/saying"
)

func main() {
	fmt.Print(saying.Greet("Pleum"))
}
