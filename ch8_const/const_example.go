package ch8_const

import (
	"fmt"
)

//--------------------------------------------------------------
// 타입 지정을 안하면 정수는 기본적으로 int로, 실수는 float64로 타입지정된다.
//--------------------------------------------------------------
const Pig = 0
const Cow int = 1
const Chicken int = 2

// 제 8장 . 상수.
// 상수는 변하지 않는 값이다.
// 한 번 정해지면 절대 변하지 않는다.
// 상수는 := (선언 대입문) 으로 할당할 수 없다.

// 구조체, 배열같은 기본타입(primitive)이 아닌 타입(complex)은 상수를 사용할 수 없다.

// !!!go에서 상수로 사용될 수 있는 타입은!!!
// bool, rune(int32), 실수, 정수, 복소수, 문자열 이다.

func ConstExample() {
	const Constval = 10
	// fmt.Println(&Constval) -> !!!compile error (상수의 주소에도 접근 못한다.)!!!

	// 1. 상수에 특정 값을 할당해서 코드의 가독성을 향상시킬 수 있다.
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)

	// 2. 열거형 (iota)
	// 일반 정수형 상수선언
	const (
		Dog  int = iota // 0부터 시작한다.
		Cat             // 1증가. -> 1
		Fish            // 1증가. -> 2
	)

	// 비트플래그 상수선언
	const (
		Bitflag1 uint = 1 << iota // 1 << 0 -> 1
		Bitflag2                  // 1 << 1 -> 2
		Bitflag3                  // 1 << 2 -> 4 (타입 미지정시 맨 처음의 타입과 동일함. = uint)
	)

	// 3. 타입 없는 상수 (연산 시 변수에 값이 복사될 때 타입이 정해지므로 여러 타입에 사용되는 상수가 필요할 때 유용하다.)
	const (
		PI          = 3.14
		PI2 float64 = 3.14
	)
	var result_float32 float32 = 3.14
	var result_float64 float64 = 6.28

	result_float32 += PI // PI는 타입없는 숫자 상수다. float32형 변수와 연산될 때 타입이 float32로 정해지므로 타입 에러가 발생하지 않는다.
	result_float64 += PI
	// result_float32 += PI2 는 당연히 타입 불일치로 컴파일 에러가 발생한다.

	fmt.Println("3. 타입없는 상수는 해당 타입 내(e.g., 실수/정수/...) 에서 유연하게 피연산될 수 있다.")
	fmt.Println("float32 타입 변수에 상수를 더한 결과: ", result_float32)
	fmt.Println("float64 타입 변수에 상수를 더한 결과: ", result_float64)

	// < 4. 상수와 리터럴 >
	// 리터럴이란 고정된 값, 즉 값 자체로 쓰인 문구라고 할 수 있다.

	// var myString string = "myString"
	// var myRune rune = 128
	// 위 "myString"이나 128 같은 고정된 값 자체의 문구가 리터럴이다.
	// Go언어는 상수를 이 리터럴과 동일한 취급을 한다.
	// 따라서 컴파일 될 때 상수는 리터럴로 변환되어 실행 파일에 쓰인다.

	// 상수 표현식 역시 컴파일 타임에 실제 결과값 리터럴로 변환하기 때문에 상수 표현식 계산에 CPU자원을 사용하지 않는다.
	// 즉, 아래 두 줄은 컴파일 타임때 그 아래 한 줄로 변환된다.
	// ######## 컴파일 타임 전과 후 ##########
	// !!상수는 컴파일 될 때 리터럴로 취급되어 상수 선언문 자체가 CPU 자원을 사용하지 않는다.!!
	// Before
	//const PI = 3.14
	//var number int = PI * 100

	// After
	//var number int = 314
	// ######## 컴파일 타임 전과 후 ##########

}

func PrintAnimal(animal int) {
	switch animal {
	case Pig:
		fmt.Println("돼지")
	case Cow:
		fmt.Println("소")
	case Chicken:
		fmt.Println("닭")
	default:
		fmt.Println("Error!")
	}
}
