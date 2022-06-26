package ch28_test_benchmark

import "fmt"

func Ch28_square() {
	fmt.Printf("30 * 30 = %d\n", Square(30))
}

func Square(age int64) int64 {
	return age * age
}
