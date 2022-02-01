package ch6_operator

import (
	"fmt"
	"math"
	"math/big"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b // 값 비교 (a가 b보다 1비트 작다면 a보다 1비트 큰(==b) 값 리턴.
	// 결국 1bit 작은 값(a)의 입력은 b와의 1bit 오차는 동일하다고 간주하는 것이다.
}

func Op_example() {
	// 산술 연산자
	// - 사칙연산자: +, -, *, /
	// - 비트연산자: %, &, |, ^(XOR 비트), &^(비트 클리어), <<, >>
	// >>로 채워지는 비트는 음수의 경우1 로 채워지고, 양수의 경우 0으로 채워진다. <<로 우항에 양의정수만 가능하며, 0으로 채워진다.

	// 대입 연산자
	// - ==, !=, <, >, <=, >=

	// 논리 연산자
	// 비교 연산자
	// 연산자 우선순위

	// [참고] int와 int64는 64비트 운영체제라도 프로그램이 다르게 인식한다..
	// 아래 연산은 컴파일 에러 발생시킨다.
	// ch6_operator\operator.go:22:6: cannot use intTest2 + intTest1 (type int64) as type int in assignment

	//var intTest1 int64 = 4611686018427387903
	//var intTest2 int64 = 4611686018427387904
	//var result int = intTest2 + intTest1
	//
	//fmt.Println(result)

	//2. 실수 연산에서 오차를 없애는 방법
	var f1, f2, f3 float64 = 0.1, 0.2, .3
	fmt.Printf("%0.18f == %0.18f : %v\n", (f1 + f2), f3, equal((f1+f2), f3))

	// 금융 프로그램이라 더 정확한 수치 계산을 위해서는 아래처럼 math/big 라이브러리의 Float객체를 사용해야 한다.
	fl1, _ := new(big.Float).SetString("0.1")
	fl2, _ := new(big.Float).SetString("0.2")
	fl3, _ := new(big.Float).SetString("0.3")

	fl4 := new(big.Float).Add(fl1, fl2)
	fmt.Println(fl1, fl2, fl3, fl4)

	// -1 리턴: fl3가 작은경우. 1은 fl4가 큰 경우. 0은 두 값이 같은 경우다.
	switch fl3.Cmp(fl4) {
	case 0:
		fmt.Println(fl3, "와 ", fl4, "는 같다.")
	case 1:
		fmt.Println(fl4, "가 ", fl3, "보다 크다.")
	case -1:
		fmt.Println(fl3, "가 ", fl4, "보다 작다.")
	default:
		fmt.Println("잘못된 입력")
	}

	// 3. 대입 연산자.
	// golang은 값 치환을 아래처럼 할 수 있다.
	fl1, fl2 = fl2, fl1

	// 정리.
	// 산술 연산자는 사칙연산과 비트연산 (특히 ^(XOR), &^(bit clear)), 시프트 연산이 있다.
	// 정수 타입으로 계산을 할 때는 해당 타입의 오버/언어 플로를 염두해야 한다.
	// 실수 타입은 math/big의 Float객체를 사용하는 것이 높은 정확도의 계산을 위해서 올바르다.
	// 복합 대입연산을 통해 한 줄로 간편하게 여러 변수에 값을 할당할 수 있다.
}
