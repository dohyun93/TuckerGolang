package main

import (
	"fmt"
)

func Divide(a, b float64) (result float64, valid bool) {
	if b == 0 {
		fmt.Println("fatal! you cannot divide with zero")
		result = 0
		valid = false
		return
	}
	result = a / b
	valid = true
	return
}

func main() {
	var a, b float64 = 7.2, 0
	result, err := Divide(a, b)
	fmt.Println("result, err: ", result, err)
}
