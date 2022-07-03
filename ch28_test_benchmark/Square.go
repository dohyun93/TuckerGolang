package ch28_test_benchmark

import (
	"fmt"
)

func Ch28_square() {
	fmt.Printf("30 * 30 = %d\n", Square(30))
	myString := "hello"
	addressMyString := &myString
	fmt.Println(myString, *addressMyString)
}

func Square(age int) int {
	return age * age
}
