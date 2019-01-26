package saying

//! Both Tests(TestGreet, ExampleGreet)
//! Benchmark | go test -bench .
import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Pleum")
	if s != "Welcome Pleum" {
		t.Error("got", s, "expected", "Welcome Pleum")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Pleum"))
	//Output
	//Welcome Pleum
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Pleum")
	}
}
