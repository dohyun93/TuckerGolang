package ch28_test_benchmark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	rst := Square(9)
	if rst != 81 {
		t.Errorf("Square(30) should be 900 but Square(30) returns %d", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := Square(3)
	if rst != 9 {
		t.Errorf("Square(3) should be 9 but returned %d", rst)
	}
}

func TestSquare3(t *testing.T) {
	// 1. Test 객체 생성
	assert := assert.New(t)
	// 2. Test 함수 호출
	assert.Equal(81, Square(9), "square(9) should be 81.")

}

func TestSquare4(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, Square(3), "square(3) should be 9.")
}
