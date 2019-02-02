package mymath

import (
	"fmt"
	"testing"
)

type blueprint struct {
	data   []int
	answer float64
}

func TestCenteredAvg(t *testing.T) {
	box := []blueprint{
		blueprint{[]int{1, 4, 6, 8, 100}, 6},
		blueprint{[]int{0, 8, 10, 1000}, 9},
		blueprint{[]int{9000, 4, 10, 8, 6, 12}, 9},
		blueprint{[]int{123, 744, 140, 200}, 170},
	}

	for _, v := range box {
		result := CenteredAvg(v.data)
		if result != v.answer {
			t.Error("Expected:", v.answer, "Got->", result)
		}
	}
}

func ExampleCenteredAvg() {
	box := []blueprint{
		blueprint{[]int{1, 4, 6, 8, 100}, 6},
		blueprint{[]int{0, 8, 10, 1000}, 9},
		blueprint{[]int{9000, 4, 10, 8, 6, 12}, 9},
		blueprint{[]int{123, 744, 140, 200}, 170},
	}

	for _, v := range box {
		fmt.Println(CenteredAvg(v.data))
		//Output:
		//6
		//9
		//9
		//170
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	solo := []int{6, 2, 15, 8, 5}
	for i := 0; i < b.N; i++ {
		CenteredAvg(solo)
	}
}
