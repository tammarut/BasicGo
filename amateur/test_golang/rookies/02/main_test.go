package main

//! ลองท่าใหม่ๆ
import "testing"

func TestSumint3(t *testing.T) {
	if sumint(1, 2) != 1+2 {
		t.Error("Expected:", 3, "Got:", sumint(1, 2))
	}
}
