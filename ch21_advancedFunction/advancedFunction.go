package ch21_advancedFunction

import (
	"fmt"
)

// ... 키워드를 사용해서 가변 인수 타입을 선언한다.
func AddNumbers(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	// nums는 int 슬라이스 타입 []int로 처리된다.
	// 즉, 가변 인수는 함수 내부에서 해당 타입의 슬라이스로 처리된다!
	for _, v := range nums {
		sum += v
	}
	return sum
}

func PrintVariousType(params ...interface{}) {
	// switch문에서 interface{}타입을 .(type)으로 각 데이터타입 케이스에 대한 처리가 가능하다.
	// 만약 해당케이스가 없더라도 default에서 처리해줄 수 있으므로, 또 변수로서 사용되는 일반 문이 아니기 때문에 패닉은 걱정하지 않아도 된다.

	// switch가 아니라면 일반적으로 if assertedValue, ok := param.(int); ok 같이 사용되기도 한다.
	// type assertion vs type conversion
	// 쉽게 말해서 런타임 vs 컴파일 타임이다.

	for _, param := range params {
		switch valType := param.(type) { // valType은 interface{}형. 현재 param의 타입을 확인한다.
		case bool:
			// type assertion.
			// interface{}타입이던 param을 bool타입으로 변환, 그 값을 val에 할당
			val := param.(bool)
			fmt.Println("type is bool. : ", val)
		case int:
			val := param.(int)
			fmt.Println("type is int. : ", val)
		case string:
			val := param.(string)
			fmt.Println("type is string. : ", val)
		default:
			fmt.Println("Unkwon type's value: ", valType)
		}
	}
}

func AdvancedFunction() {
	fmt.Println()
	// 함수 고급편.
	// [1] : 가변 인수 함수 -> 인수 개수를 고정하지 않고 받을 수 있다.
	// [2] : defer 지연 실행 -> 함수 종료 전에 반드시 실행해야 하는 코드
	// [3] : 함수타입 변수 -> 함수를 값으로 갖는 타입이다.
	// [4] : 함수 리터럴(익명함수 또는 람다 라고도함) -> 이름 없는 함수를 정의하고 함수 타입 변수에 대입할 수 있다.

	// # [1] : 가변 인수 함수 # -> 키워드 '...'을 사용한다!
	fmt.Println(AddNumbers(1, 2, 3, 4, 5))
	fmt.Println(AddNumbers(10, 20))
	fmt.Println(AddNumbers())

	// int말고 빈 인터페이스 interface{}로 가변 인자들을 받아보자
	PrintVariousType(2, 3.28, true, "stringVal")
	// # [2] : defer 지연 실행 #

	// # [3] : 함수타입 변수 #

	// # [4] : 함수 리터럴 == 람다 == 익명 함수
}
