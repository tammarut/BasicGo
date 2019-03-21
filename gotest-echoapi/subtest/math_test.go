package subtest

import "testing"

func TestMath(t *testing.T) {
	m := Math{
		A: 5,
		B: 10,
	}

	t.Run("Plus", func(t *testing.T) {
		result := m.Plus()
		expected := 15

		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})

	t.Run("Multiply", func(t *testing.T) {
		result := m.Multiply()
		expected := 50

		if result != expected {
			t.Errorf("Expected: %v, Got: %v", expected, result)
		}
	})

}
