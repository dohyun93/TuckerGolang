package main

import (
	"fmt"
)

func main() {
	// 1. 연산자 연습
	// +, -, *, /, % -> 산술연산
	// &, ^, |, &^ -> 비트연산
	// <<, >> -> 시프트 연산

	// 비트연산 - &^
	// &^ 연산자는 bit clearing 연산자이다.
	fmt.Println("1. 비트연산자 - &^ (bit clearing)")
	var mybit int = 38   // 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00100110
	var clearing int = 6 // 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000110
	fmt.Println("before clearing, mybit and clearing are: ", mybit, clearing)
	mybit &^= clearing // clearing 이 1인 비트들에 대해 mybit의 동일 비트를 0으로 클리어링 하는 연산.
	fmt.Println("after clearing, mybit is: ", mybit)
	fmt.Println()

	// ^ (XOR) 연산자
	fmt.Println("2. 비트연산자 - ^ (XOR)")
	var mybit2 int8 = 4
	var mybit3 int8 = 2
	fmt.Println("00000100 을 00000010 과 XOR 연산한다.")
	mybit2 ^= mybit3
	fmt.Println("결과는: ", mybit2) // 000000110 -> 6
	fmt.Println()

	// shift 연산자 (오른쪽은 0~양의정수만 올수있다.)
	fmt.Println("3. 시프트연산자 - <<")
	var myInt int = 32
	fmt.Println("왼쪽으로 비트연산하기 전: ", myInt)
	myInt = myInt << 1
	fmt.Println("왼쪽으로 << 1 연산 후: ", myInt)
}
