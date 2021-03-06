package main

//!Data are slice <-func gen()
//! main call func CenteredAvg =>find the average of a list of numbers

import (
	"fmt"

	"github.com/tammarut/amateur/ninja_lv13/03/mymath"
)

func main() {
	xxi := gen() //=> xxi kept func gen() that return e
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d} //=> partie a,b,c,d
	return e
}
