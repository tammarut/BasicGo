package word

import (
	"fmt"
	"testing"
)

const message = `we are hero` //=> mockup data

//* Regular Test(func Count)
func TestCount(t *testing.T) {
	result := Count(message)
	if result != 3 {
		t.Error("Expected: 3   Got:", result)
	}
}

//* Challenge Testcase(func UseCount)
func TestUseCount(t *testing.T) {
	result := UseCount(message)
	for k, v := range result { //? What are these code..?
		switch k {
		case "we":
			if v != 1 {
				t.Error("Expected: 1   Got:", v)
			}
		case "are":
			if v != 1 {
				t.Error("Expected: 1   Got:", v)
			}
		case "hero.":
			if v != 3 {
				t.Error("Expected: 3  Got:", v)
			}
		}
	}
}

//* Example test(func Count)
func ExampleCount() {
	fmt.Println(Count(message))
	//Output:
	//3
}

//*Benchmark (func Count)
func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(message)
	}
}

//* Bencgmark (func UseCount)
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(message)
	}
}
