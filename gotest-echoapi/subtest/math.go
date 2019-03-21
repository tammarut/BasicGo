package subtest

type Math struct {
	A int
	B int
}

func (m Math) Plus() int {
	return m.A + m.B
}

func (m Math) Multiply() int {
	return m.A * m.B
}
