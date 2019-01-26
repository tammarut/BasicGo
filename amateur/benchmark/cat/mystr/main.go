package mystr

//! Cat = manual forloop
//! Join = standard lib
//* Go Definition <Join> to learn
import "strings"

func Cat(slicecut []string) string {
	s := slicecut[0]

	for _, v := range slicecut[1:] {
		s += " "
		s += v
	}
	return s
}

func Join(slicecut []string) string {
	return strings.Join(slicecut, " ")
}
