package ch21_advancedFunction

import (
	"fmt"
	"os"
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
	// 함수가 종료되기 직전에 실행해야 하는 코드는 defer로 처리할 수 있다

	// 예를 들어 파일이나 소켓 핸들처럼 OS 내부 자원을 사용하는 경우 파일을 생성하거나 읽을 때 OS에 파일 핸들을 요청한다.
	// 그러면 OS는 파일 핸들을 만들어서 프로그램에 알려준다.
	// 하지만 이런 자원은 OS 내부 자원이기 때문에 사용하고 반드시 OS에 되돌려줘야 한다. 프로그램에서 OS 내부 자원을 되돌려주지 않으면
	// 내부 자원이 고갈되어 더는 파일을 만들지 못하거나 네트워크 통신을 하지 못할 수 있다.

	// 즉, 파일 핸들을 반환하는 것 과 같이 함수 종료전 반드시 실행되어야 하는 기능은 defer로 정의해 주는 것이 좋다.
	// defer를 이용한 '지연 실행'!!

	f, err := os.Create("ch21_advancedFunction.txt")
	if err != nil {
		fmt.Println("Failed to create a file!")
		return
	}

	defer fmt.Println("3. 반드시 파일 핸들러 자원을 OS에 반환하세요.") // 3
	defer f.Close()
	defer fmt.Println("2. 파일을 닫습니다.") // 2

	fmt.Println("1. 파일에 Hello World!를 씁니다.") // 1
	fmt.Fprintln(f, "Hellow World!")
	// defer은 스택처럼 역순으로 호출된다!!!!
}

////////// step3 function type variable /////////////
// # [3] : 함수타입 변수 #
// 함수타입 변수란 함수를 값으로 갖는 변수를 의미한다.

// 숫자값을 가지는 변수, 그리고 메모리 주소를 값으로 갖는 포인터처럼 함수의 시작점도 숫자로 나타낼 수 있다.
// 프로그램 카운터는 CPU내 레지스터에 위치하는 '다음 실행될 프로그램의 라인을 지정'한다.
// 다음에 실행될 함수의 라인을 가리키는 함수 포인터를 타입으로 갖는, 함수 타입 변수를 알아보자.

func Step3_FunctionTypeVariable() {
	functionTypeVariable()
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	switch op {
	case "+":
		return add
	case "*":
		return mul
	default:
		fmt.Println("Unknown operator")
		return nil
	}
}

func functionTypeVariable() {
	// 정수 두 개를 입력받아 정수 하나를 리턴하는 함수 타입 변수 선언.
	var myFunction func(int, int) int
	myFunction = getOperator("*")

	var result = myFunction(10, 20)
	fmt.Println(result)
}

// 참고로 func (int, int) int 는 타입으로 쓰기 너무 길다.
// 이럴 경우 앨리어싱 으로 별칭을 부여해 사용하는 것이 효율적일 때가 있다.

////////// step 4 function literal ////////////
// # [4] : 함수 리터럴 == 람다 == 익명 함수
// 함수 리터럴은 익명함수, 또는 람다로 불린다.
// 함수명을 적지 않고, 함수 타입 변수로 대입되는 함수값을 의미.

// 함수 리터럴을 사용해서 익명함수로 리턴하는 구현을 할 수 있다.
type opFunc func(a, b int) int

func getOpFunction(operator string) opFunc {
	switch operator {
	case "+":
		return func(a, b int) int { // 함수 별칭 사용안됨.
			return a + b
		}
	case "*":
		return func(a, b int) int {
			return a * b
		}
	default:
		return nil
	}
}

func Step4_FunctionTypeVariable() {
	// myFunc := getOpFunction("+")
	// result := myFunc(10, 20)

	// 위 두 줄을 아래처럼 할 수도 있다.
	result := func(a, b int) int {
		return a + b
	}(10, 20)

	fmt.Println("연산결과는: ", result)
}

// 4-1. 함수 리터럴 내부 상태
// 함수 리터럴은 필요한 변수를 내부 상태로 가질 수 있다.
// 함수 리터럴 내부에서 사용되는 외부 변수는 자동으로 함수 내부 상태로 저장된다.

func Step4_1_Function_Literal_InternalState() {
	i := 0        // 1. 함수 리터럴 외부 변수 i값 선언.
	f := func() { // 2. 함수 리터럴 정의.
		i += 10
	}
	i++ // 3. i값 +1
	f() // 4. 함수 리터럴은 값 복사가 아닌 인스턴스 참조로 가져오기 때문에, 함수 리터럴에서 외부 변수 덧셈을 하면 외부변수 i의 값이 바뀐다.!

	// * 추가 * '캡쳐'란?
	// '캡쳐'는 '함수 리터럴 외부 변수(여기서 i)'를 함수 리터럴 '내부 상태로 가져오는 것'을 말한다.
	// 캡쳐는 값 복사가 아닌 참조 형태로 가져온다.!!

	// 즉, 외부 변수의 주소값을 포인터 형태로 함수 리터럴 내부 상태로 가져온 뒤 (캡쳐),
	// 나중에 캡쳐된 내부 상태를 사용할 때 외부 변수의 메모리 주소값을 통해 외부 변수에 접근한다.
	fmt.Println("i: ", i)
}

func captureLoop() {
	f := make([]func(), 3) // 함수 리터럴 타입 슬라이스 (길이: 3) 선언
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	// f[0]~f[2]는 func(){fmt.Println(i)} 함수 리터럴이다.
	// 여기서 외부변수 i는 참조로 캡쳐되기 때문에 i의 최종 값 3을 갖는 인스턴스를 참조한다.
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func captureLoop2() {
	f := make([]func(), 3)
	fmt.Println("Valueloop2")
	for i := 0; i < 3; i++ {
		v := i // v 변수에 i값 복사
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func Capture_FunctionLiteral() {
	captureLoop()
	captureLoop2()
}

//////////////// 파일 핸들을 내부 상태로 사용하는 예 //////////////////
// 함수 리터럴을 이용해서 원하는 함수를 그때그때 정의해서 함수 타입 변수값으로 사용할 수 있다.
// 그리고 필요한 외부 변수를 내부 상태로 가져와서 편리하게 사용할 수 있다.

// 파일 핸들을 내부 상태로 갖는 예를 살펴보자.

type Writer func(string)

func WriteHello(writer Writer) { // 함수 타입 변수 writer 호출.
	writer("Hello Golang!") // msg.
}

func FileHandleInnerState() {
	f, err := os.Create("ch21_functionLiteral.txt")
	if err != nil {
		fmt.Println("Failed to create the file")
		return
	}
	defer f.Close()

	WriteHello(func(msg string) {
		fmt.Fprintln(f, msg) // 함수 리터럴에서 외부 변수 f 사용.
	})
}

// 21장 함수 고급편 정리
// 1. 인수 타입 앞에 ...을 찍으면 해당 타입의 가변길이 슬라이스 인수를 표현한다.
// 2. defer (지연실행)을 통해 함수 종료 전 반드시 처리할 로직을 수행한다. (단 선언 역순으로 호출됨.)
// 3. 함수 타입 변수를 사용해서 함수를 변수의 값으로 저장할 수 있다.
// 4. 함수 리터럴을 사용해서 간편하게 내부 상태를 갖는 함수를 정의할 수 있다.
