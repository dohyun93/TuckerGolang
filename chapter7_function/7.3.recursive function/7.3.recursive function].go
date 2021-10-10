package main

import (
	"fmt"
)

func fibo_recurFunc(number int64) (result int64) {
	// 재귀함수 상단부에 반드시 종료조건 구현해주는 것 필수.
	if number == 1 {
		return 1
	} else {
		return fibo_recurFunc(number-1) * number
	}
}

func main() {
	var myNum int64 = 3
	result := fibo_recurFunc(myNum)
	fmt.Printf("The fibonacci result of %d is %d", myNum, result)
}
