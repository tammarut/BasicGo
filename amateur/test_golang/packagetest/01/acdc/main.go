package acdc

//! Regular func sum variadic
func Sumint(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
