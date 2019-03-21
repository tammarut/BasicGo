package subtest

//Math is struct...
type Math struct {
	A int
	B int
}

//Plus is regular plus...
func (m Math) Plus() int {
	return m.A + m.B
}

//Multiply is ...
func (m Math) Multiply() int {
	return m.A * m.B
}
