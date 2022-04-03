package ch4_Variable

import (
	"fmt"
)

// 함수 바깥에서 대문자로 시작하는 변수는 패키지 외부에서 참조가 가능하다.
// var AmIVisible = 3

// 4장 (변수) 핵심 요약
// 1. 변수는 값을 저장하는 메모리의 공간으로, 변수 이름이 곧 그 메모리의 시작 주소를 의미한다.
//  개발자는 이 변수명(메모리 시작주소)으로 그 메모리의 값에 쉽게 접근할 수 있다.
//
// 2.변수 선언은 컴퓨터에게 값을 저장할 공간을 메모리에 마련하라고 명령을 내리는 것이다. 이를 메모리 할당이라고 한다.
// 3. 변수는 4가지 속성을 지닌다. 이름, 값, 타입, 주소.
// 4. 변수 선언은 var, 타입을 명시하거나 선언 대입문(:=)을 사용하여 생략하기도 한다.
// 5. 타입 변환은 한 타입의 값을 다른 타입으로 변환시키는 것이다. Go언어는 최강타입이기 때문에 같은 정수라도 타입이 다르면 연산이 안된다.
//    또한 자동 형 변환이 없기 때문에, 다른 타입의 변수들을 연산을 할 경우 명시적 형 변환을 해주어야 한다.

func VariableType() {
	fmt.Println("1. unsigned integer")
	fmt.Println("uint8, uint16, uint32, uint64")

	fmt.Println("2. signed integer")
	fmt.Println("int8, int16, int32, int64")

	fmt.Println("3. float")
	fmt.Println("float32, float64")

	fmt.Println("4. Complex number")
	fmt.Println("complex64, complex128 - 진수, 가수는 각각 float32 범위와 같음")

	fmt.Println("5. etc")
	fmt.Println("byte는 uint8의 별칭.")
	fmt.Println("rune은 UTF-8로 문자하나 표현가능(int32의 별칭)")
	fmt.Println("int는 32비트 CPU 컴퓨터는 int32, 64비트 CPU 컴퓨터는 int64와 동일")
	fmt.Println("uint는 32비트 CPU 컴퓨터는 uint32, 64비트 CPU 컴퓨터는 uint64와 동일")
}

func VariableSum(a, b int) (result int) {
	result = a + b
	fmt.Println(a, " + ", b, " = ", result)
	return result
}

func MemAllocAndInitialization(a int, b string) {
	// 변수는 컴퓨터에게 데이터를 저장할 수 있는 공간이다.
	// 변수를 선언함으로써 메모리를 할당하고, 그 변수에 값을 할당할 수 있다.
	var intValue int = 100 // 컴퓨터가 이 라인을 실행할 때, IntValue 라는 이름의 정수형 변수에 정수형만큼의 메모리를 할당한다.
	// 그리고, 그 메모리 공간의 시작 주소를 IntValue라고 지칭한다.
	// 즉, 변수의 이름은 그 변수에 담긴 데이터를 저장하는 메모리의 시작 주소이며, 변수 타입크기만큼 메모리가 할당된다.

	var myString string = "I'm good at Golang."

	a = intValue
	b = myString

	fmt.Println("내 자산은 ", a, "억원이고, ", b)

	// 변수 이름: 개발자는 변수 이름을 통해 값이 저장된 메모리의 시작 주소를 접근할 수 있다.
	// 값: 변수가 가리키는 메모리 공간에 저장된 값이다.
	// 주소: 변수의 값이 저장된 메모리 공간의 시작 주소다.
	// 타입: 변숫값의 형태다. 정수/실수/문자열 등 다양한 타입이 있고, 데이터 타입이라고도 한다.

	// 만약 함수 바깥에 대문자로 시작하는 변수는 패키지 외부로 공개된다.
}

func TypeDefaultValue() {
	// 정수형 변수들의 기본 할당값은 0이다.
	// uint8, uint16, uint32, uint64, int8, int16, int32, int64, int, uint, byte, rune

	// 실수형 변수들의 기본 할당값은 0.0이다.
	// float32, float64, complex64, complex128

	// bool 타입 기본값은 false 이다.

	// 문자열 타입 기본값은 ""이다.

	// 그 외
	// 정의되지 않은 메모리 주소를 나타내는 Go 키워드인 nil로 정의된다.

	//--------------------------------------------------------------
	// 타입 지정을 안하면 정수는 기본적으로 int로, 실수는 float64로 타입지정된다.
	//--------------------------------------------------------------
	// 선언 대입문 (:=)
	// 선언 대입문은 var와 타입을 생략하여 값을 변수에 할당할 수 있다.

	// Go언어는 최강 타입 언어라서 자동 형 변환이나 같은 정수라도 타입이 다르면 연산이 안된다.

	myInt := 30        // int
	mySalary := 7120.9 // float64

	someValue := int(mySalary) + myInt // 기본적으로 최강타입언어라 동일 타입끼리만 연산 가능함. 같은 정수라 할지라도.
	fmt.Println("someValue", someValue)

	// 그리고 int16에서 int8로 형변환 하면 상위 8비트가 날라감에 주의.

	//--------------------------------------------------------------
	// 소수 표현은 소수부와 지수부로 나뉠 수 있다.
	// 0.12345e+05 에서 소수부는 12345 이고 지수부는 5이다.
	// 소수 표현범위는 소수부의 자리(비트)수가 많을수록 더 넓어진다.
}

func Int16ToInt8() {
	// cast vs conversion?

	var Int16 int16 = 1409 // 00000101 10000001 -> int8로 변경 시 좌측의 상위 8bit가 truncate(잘림)됨. signed 데이터타입일 경우
	// 음수의 값으로 변경됨 (10000001은 -127임. 즉 int16에서 int8로 변경 시 1409에서 -127이 됨.)
	var Int8 int8 = 1

	fmt.Printf("\n\nInt16 is: %d, and Int8 is: %d\n", Int16, Int8)
	fmt.Printf("Int16(Type, value: %T, %v)\n", Int16, Int16)

	fmt.Printf("Now, Let's change the Int16 to Int8 and see what happens\n")
	fmt.Printf("int8(Int16): %d", int8(Int16))

	var Int8_2 int8 = -127 // 10000001
	var Int16_2 int16      // sign extension. 11111111 10000001

	fmt.Printf("Int8_2 : %d\n", Int8_2)
	Int16_2 = int16(Int8_2)
	fmt.Printf("Int8_2 to int16: %d\n", Int16_2) // Int16_2: -127
}
