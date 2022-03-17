package ch23_errorHandling

// panic은 프로그램을 정상 진행시키기 어려운 상황을 만났을 때 프로그램 흐름을 중지시키는 기능이다.
// Go 언어는 내장함수 panic()으로 패닉 기능을 제공한다.
// 지금까지 error 인터페이스를 사용해 에러 처리를 했다.

// error 인터페이스를 사용하면 호출자에게 에러가 발생한 이유를 알려줄 수 있었다.
// 그런데 예기치 못한 에러에 직면하는 경우가 있다. 예를 들어 잘못된 메모리에 접근하거나 메모리가 부족하면 프로그램이 실행 불가능 할 수 있다.

// 이럴때는 프로그램을 강제 종료해서 문제를 빠르게 파악하는 편이 나을 수 있다.
// panic() 함수를 사용하면 문제 발생 시점에 프로그램을 바로 종료시켜서 빠르게 문제 발생 시점을 알 수 있다. (버그수정에 유용한 방식)

// panic()을 호출하고 인수로 에러 메시지를 입력하면 프로그램을 즉시 종료하고 에러 메시지를 출력 + 함수 호출 순서를 나타내는 콜 스택을
// 표시한다. 이 정보를 사용해 에러가 발생한 경로를 파악할 수 있다.

// 나눗셈에서 제수가 0이면 panic()을 호출하는 예제를 살펴보자.

import (
	"fmt"
)

func Divide(a, b int) {
	if b == 0 {
		panic("0으로 나눌 수 없다.") // func panic(v interface{})
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func PanicExample1() {
	Divide(9, 3)
	Divide(9, 0)
	// 아래처럼 panic이 발생하면 콜스택을 보여준다.
	// 콜 스택은 첫 호출부터 panic이 발생한 마지막 함수까지 호출된 순서의 역순을 담고있다.
	// panic 발생시 panic이 발생한 위치부터 알려준다.

	//	goroutine 1 [running]:
	//	TuckerGolang/ch23_errorHandling.Divide(0x9, 0x0)
	//		C:/Go/src/TuckerGolang/ch23_errorHandling/5. panic.go:24 +0x175
	//	TuckerGolang/ch23_errorHandling.PanicExample1()
	//		C:/Go/src/TuckerGolang/ch23_errorHandling/5. panic.go:31 +0x50
	//	main.chapter23(...)
	//		C:/Go/src/TuckerGolang/main.go:57
	//	main.main()
	//		C:/Go/src/TuckerGolang/main.go:49 +0x27
}

// 그런데 만약 panic이 발생한다면 그냥 실행중이던 프로그램이 죽기때문에 에러 메시지를 표시하고 복구를 시도하는게 더 나을 수도 있다.
// 사용자에게 프로그램이 배포된 이후에는 복구할 수 있는 패닉이라면 복구를 하는게 낫다.

// 2. 패닉 전파 및 복구
// panic은 호출 순서를 거슬러 올라가며 전파된다.
// 만약 main->f()->g()->h() 함수 호출과정 중 h()에서 panic이 발생하면 panic은 역순인 h()->g()->f()->main 으로 전파된다.

// 만약, main() 함수에서까지 복구되지 않으면 프로그램이 강제 종료된다.
// 어느 단계에서든 패닉은 복구된 시점부터 프로그램이 계속된다. recover()이라는 함수로 panic 복구를 할 수 있는데, 예시를 살펴보자.
func f() {
	fmt.Println("f() 함수 시작")
	// 4. panic 복구
	// panic이 f까지 전파되었으나, defer를 사용해 함수 종료 전 함수 리터럴이 실행된다.
	// 함수 리터럴 내부에서 recover()를 사용해 패닉 복구를 시도한다.
	// 전파중인 panic이 있으므로 복구가 되고 panic 메시지(r)를 출력한다...!!!
	// recover()는 발생한 panic 객체를 반환해준다...
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g() // 1. f() -> g() -> h() 순서로 호출
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0)) // 2. h() 함수 호출 - 패닉.
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없다.") // 3. panic 발생!
	}
	return a / b
}

func PanicExample2() {
	f()
	fmt.Println("프로그램이 계속 실행된다.")
}

// recover()는 제한적으로 사용하는게 좋다.
// panic이 발생하면 그 즉시 호출 순서 역순으로 전파하기 때문에 복구가 되더라도 프로그램이 불안정할 수 있다.
// 예를 들어 파일에 데이터를 쓰는 프로그램에서 데이터를 일부만 쓴 상태에서 패닉이 발생하고 다시 복구하면 데이터가 비정상적으로 저장된 상태로
// 남게 된다. 이럴 때는 그냥 복구하지 않거나 데이터가 비정상적으로 남아있지 않도록 확실히 지워줘야 한다.

// recover은 interface{} 타입을 반환한다.
// 그렇기 때문에, 실제 recover()로 반환된 타입을 사용하려면 type assertion을 해야 한다.
// 예시
// if r, ok := recover().(net.Error); ok{
// 	fmt.Println("r is net.Error Type")
// }
