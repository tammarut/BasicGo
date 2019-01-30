package dog

//! Hightligt! we come for this.
//! consists of [Test, Example, Benchmark]
import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	result := Years(10)
	if result != 70 {
		t.Error("Expected: 70 Got:", result)
	}
}

func TestYesrsTwo(t *testing.T) {
	result := YearsTwo(10)
	if result != 70 {
		t.Error("Expected:", 70, "Got:", result)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	//Output:
	//70
}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//Output:
	//70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
