package ch9_if

import (
	"fmt"
)

func testHeight(height int) (old bool) {
	fmt.Println("Height 측정")
	if height > 180 {
		old = true
		return
	} else {
		old = false
		return
	}
}

func getAge() (age int, err error) {
	var Age int

	n, Err := fmt.Scanln(&Age)
	if Err != nil {
		fmt.Println("Scanln 오류발생!")
		age = -1
		err = Err
	} else {
		fmt.Println(n, "개의 변수 입력!")
		age = Age
		err = nil
	}
	return
}

func IfExample() {
	var currentSign, nextSign string = "Red", "Green"
	if currentSign == "Green" {
		fmt.Println("현재 건너도 됩니다.")
	} else {
		fmt.Println("다음 신호때 건너세요.")
	}

	if nextSign == "Red" {
		fmt.Println("현재 건너도 됩니다.")
	} else {
		fmt.Println("다음 신호때 건너세요.")
	}

	// 1. 쇼트 서킷 !!!
	// 조건문에서 &&의 경우 좌변에 false가 있다면 그 우변의 항을 계산하지 않고 false처리를 한다.
	// 비슷하게 ||의 경우 좌변에 true가 있다면 그 우변의 항을 계산하지 않고 true처리를 한다.
	// 조건문 우변에 있는 코드가 실행되지 않을 수 있으니 이를 염두하여 개발해야 한다.
	var myAge int = 30
	var myHeight int = 187
	if myAge > 999999 && testHeight(myHeight) { // 좌변 거짓 && 우변 --> 우변 testHeight 함수 호출 안함.
		fmt.Println("조건 만족")
	} else {
		fmt.Println("조건 불만족")
	}

	if myAge > -1 || testHeight(myHeight) { // 좌변 참 || 우변 --> 우변 testHeight 함수 호출 안함.
		fmt.Println("조건 만족")
	} else {
		fmt.Println("조건 불만족")
	}

	// 2. if 초기문; 조건문
	// 초기문에 선언된 변수는 if~else if~else 문으로 한정된다는 것을 유의하자.!!
	if Age, Err := getAge(); Err != nil {
		fmt.Println("getAge() 함수 호출 오류 발생!! Age: ", Age)
	} else {
		fmt.Println("getAge() 함수 호출 성공! Age: ", Age)
	}
}
