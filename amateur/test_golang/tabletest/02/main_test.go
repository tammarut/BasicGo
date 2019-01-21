package main

//! Testfunction => struct+slice+forloop(check)
import "testing"

func TestSumint(t *testing.T) {
	type blueprint struct {
		data   []int
		answer int
	}

	box := []blueprint{
		blueprint{[]int{2, 3, 4}, 9},
		blueprint{[]int{42, 43, 44}, 129},
		blueprint{[]int{5, 7}, 12},
		blueprint{[]int{7, 7, 7}, 21},
		blueprint{[]int{14, 18}, 32},
		blueprint{[]int{12, 19}, 31},
		blueprint{[]int{1, 1}, 2},
	}

	for _, v := range box {
		result := sumint(v.data...)
		if result != v.answer {
			t.Error("Expected:", v.answer, "Got->", result)
		}
	}
}
