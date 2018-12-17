package main

import "testing"

//commit test again
func TestPlus(t *testing.T) {
	result := plus(2, 3)

	if result != 5 {
		t.Errorf("Expected result 5, but it was %v", result)
	}
}

func TestPrism(t *testing.T) {
	result := prism(2, 3, 2)

	if result != 12 {
		t.Errorf("Expected result 12, but it was %v", result)
	}
}

func BenchmarkFact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fact(5)
	}
}
func BenchmarkFactForloop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factForloop(5)
	}
}

//2 tests below not done yet.
/*func TestSwap(t *testing.T) {
	result := swap(0, 1)

	if result != 1,0 {
		t.Errorf("Expected 1,0, but it was %v", result)
	}
}

func TestFindFirstAndLast(t *testing.T){

} */
