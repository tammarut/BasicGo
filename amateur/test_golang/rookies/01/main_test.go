package main

//! They're test cases.
import "testing"

func TestSumint3(t *testing.T) {
	result := sumint(1, 2)
	if result != 3 {
		t.Error("Expected:", 3, "Got->", result)
	}
}

func TestSumint6(t *testing.T) {
	result := sumint(1, 5)
	if result != 6 {
		t.Error("Expected:", 6, "Got->", result)
	}
}

func TestSumint10(t *testing.T) {
	result := sumint(1, 9)
	if result != 10 {
		t.Error("Expected:", 10, "Got:->", result)
	}
}
