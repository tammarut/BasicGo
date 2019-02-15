package word

//! 2 func
//!   1) return map[string]int
//!   2) return message's length
import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s) //=> Just WOW!: Split by space" "
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m //=> return map "key": "value"
}

func Count(s string) int {
	xs := strings.Fields(s) //=> Just WOW!: Split by space" "
	return len(xs)          //=> return length slice(xs)
}
