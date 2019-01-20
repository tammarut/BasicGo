//! Test pattern struct+slice
package main

import "testing"

func Testsumint(t testing.T) {
	type blueprint struct {
		data   []int
		answer int
	}

	box := []blueprint{
		blueprint{[]int{21, 21}, 42},
		blueprint{[]int{3, 4, 5}, 12},
		blueprint{[]int{1, 1}, 2},
		blueprint{[]int{-1, 0, 1}, 0},
	}

	for _, v := range box {
		result := sumint(v.data...)
		if result != v.answer {
			t.Error("Expected:", v.answer, "Got->", result)
		}
	}
}
