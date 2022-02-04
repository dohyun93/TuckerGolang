package ch7_function

import (
	"fmt"
)

// 함수는 다중 값 반환 가능....

func totalAvgScore(a, b, c int, d float64) (result float64, bigger bool) {
	totalSum := a + b + c
	result2 := float64(totalSum) / 3.0
	result2 += d
	result = result2 / 3
	if result > 10 {
		bigger = true
	} else {
		bigger = false
	}
	return result, bigger
}

func FunctionExample() {
	// 함수는 반복되는 코드를 기능으로 제공하여 인자로 다양한 동작을 가능하게 하는 기능이다.
	var dohyun, hyangwon, taeran int = 100, 99, 200
	var bias float64 = 2.7

	result, isBigger := totalAvgScore(dohyun, hyangwon, taeran, bias)

	if isBigger {
		fmt.Println(result, "is bigger than 10")
	} else {
		fmt.Println(result, "is smaller than 10")
	}
}
