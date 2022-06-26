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
