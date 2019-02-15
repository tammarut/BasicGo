package mymath

//! Test, Example, Benchmark
import (
	"fmt"
	"testing"
)

type blueprint struct {
	data   []int
	answer float64
}

//* Test ->func CenteredAvg
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

//* Example func CenteredAvg
func ExampleCenteredAvg() {
	solo := []int{1, 4, 6, 8, 100}
	fmt.Println(CenteredAvg(solo))
	//Output:
	//6

}

//* Benchmark func CenteredAvg
func BenchmarkCenteredAvg(b *testing.B) {
	solo := []int{1, 4, 6, 8, 100}
	for i := 0; i < b.N; i++ {
		CenteredAvg(solo)
	}
}
