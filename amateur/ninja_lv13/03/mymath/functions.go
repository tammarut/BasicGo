package mymath

import "sort"

//! CenteredAvg computes the average of a list of numbers
//! after removing the largest and smallest values.
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)          //=> arrange in ascending order
	xi = xi[1 : len(xi)-1] //=> select whole data Except: first, last

	n := 0 //=> total
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi)) //=> total/number
	return f
}
