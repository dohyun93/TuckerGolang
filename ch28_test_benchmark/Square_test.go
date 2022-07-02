package ch28_test_benchmark

import (
	"testing"
)

func TestSquare(t *testing.T) {
	rst := Square(30)
	if rst != 900 {
		t.Errorf("Square(30) should be 900 but Square(30) returns %d", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := Square(3)
	if rst != 9 {
		t.Errorf("Square(3) should be 9 but returned %d", rst)
	}
}
