package ch19_method

import (
	"fmt"
)

type account struct {
	balance int
}

// Go에는 클래스가 없다.
// Go에는 구조체만 있다.

// Go에는 구조체 밖에 메서드를 지정한다.
// 어느 구조체에 속하는지 메서드를 표시하기 위해 '리시버'라는 기능을 사용한다.
// 즉, 리시버는 메서드가 속하는 타입을 알려주는 기법이다.

func MethodExample() {
	fmt.Println()
	acc := &account{balance: 100}
	MinusFunction(acc, 20)
	acc.MinusMethod(20)
	fmt.Println(acc.balance)
}

// MinusReceiver 메서드는 account 구조체 타입에 속한다는걸 알 수 있다.
func (acc *account) MinusMethod(minus int) {
	acc.balance -= minus
}

func MinusFunction(acc *account, minus int) {
	acc.balance -= minus
}

// 일반적으로 리시버는 리시버 타입이 선언된 파일 안에 정의한다.
// 위 처럼 account라는 구조체가 있는 이 파일에 account 타입의 리시버들을 정의한다.

// 리시버 타입은 별칭의 타입에 속할 수 있다.
// e.g., type myInt int 일 때, myInt타입의 리시버를 정의할 수 있다.

// [------------------- 메서드는 왜 필요한가? --------------------]
// 1. 기능(behaviour) 과 데이터의 응집성을 높인다.
// -> 메서드는 데이터의 기능을 정의한다.
// -> 함수는 소속이 없지만, 메서드는 데이터에 관련된 기능을 정의하며, 해당 데이터에 대한 소속을 갖는다.
// -> 데이터와 그 기능을 묶어놓아 응집도를 높여 유지보수를 더 용이하게 하고, 따라서 산탄총 문제가 발생하지 않게 해준다.

// [------------------- 객체(타입)와 인스턴스 객체 ----------------]
// MinusMethod 메서드는 account 구조체를 리시버로 갖는다.
// 즉 이 메서드는 account 구조체에 속한다.
// 이제 account 는 데이터 뿐 아니라 기능도 추가된 '객체'가 되었다.
// 이 객체(타입)의 인스턴스를 '객체 인스턴스'라고 한다.

// 객체 간 상호관계는 메서드로 표현한다. -> OOP
// Go에는 클래스와 상속이 없고 메서드와 인터페이스만 지원한다.
// 그래서 일부는 Go가 OOP언어가 아니라고 하지만, 필자는 클래스와 상속 지원여부가 아닌 객체 간의 상호관계 중심으로 프로그래밍 가능한지 여부가
// OOP언어라고 할 수 있는 이유라고 보기 때문에, Go를 OOP언어라고 구분한다.

// account 포인터 구조체가 리시버 (같은 인스턴스를 가리킨다.)
func (ptrRec *account) MultiplyMethod(value int) {
	ptrRec.balance *= value
}

// account 구조체가 리시버 (값 복사)
func (valRec1 account) MultiplyFunction(value int) {
	valRec1.balance *= value
}
func (valRec2 account) MultiplyFunctionReturn(value int) account {
	valRec2.balance *= value
	return valRec2
}

func MethodVSFunction() {
	var ptrVar *account = &account{balance: 800}
	// ptrVar과 ptrRec는 동일한 인스턴스를 가리킨다.
	ptrVar.MultiplyMethod(2)
	fmt.Println(ptrVar.balance)

	// Go 내부적으로 (*ptrVar)로 값 타입으로 바꾼 뒤, 호출해줌.
	// ptrVar의 구조체 값만 valRec1로 복사해줌. 함수 내에서 바뀌는건 valRec1의 데이터일 뿐.
	ptrVar.MultiplyFunction(3) // valRec1 구조체에 값만 복사. -> ptrVar가 가리키는 인스턴스의 데이터 불변.
	fmt.Println(ptrVar.balance)

	// ptrVar의 값을 valRec2로 넘겨주고 리턴한 걸 valVar에 복사.
	// valVar, valRec2, ptrVar 모두 다른 메모리 주소를 갖는 인스턴스다.
	// (*ptrVar).Multi~ 로 내부적으로 값 타입으로 바꿔서 호출한다.
	var valVar account = ptrVar.MultiplyFunctionReturn(2)
	fmt.Println(valVar.balance)

	// (&valVar).Multi~ 로 내부적으로 포인터 타입으로 바꿔서 호출한다.
	// 따라서 valVar과 ptrRec 리시버 변수는 동일한 객체를 가리키는 인스턴스다.
	valVar.MultiplyMethod(2)
	fmt.Println(valVar.balance)
}

// [----------- 19장 메서드 정리 ----------]
// 포인터 메서드는 리시버 변수가 가리키는 인스턴스 중심이고,
// 값 타입 메서드는 리시버 변수의 값 중심이 된다.

// 핵심 요약
// 1. 리시버가 있으면 메서드, 없으면 함수다.
// 2. 리시버는 메서드를 호출하는 주체로, 메서드는 리시버를 통해서만 호출될 수 있다.
//    따라서 메서드는 리시버에 속한 기능을 표현한다. 모든 로컬 타입은 리시버가 될 수 있다.
// 3. 메서드를 통해서 데이터와 기능이 묶임으로써 객체라는 개념이 생겼고, 프로그래밍 패러다임은
//    절차지향에서 객체지향으로 변화했다.
// 4. 포인터 메서드는 인스턴스 중심으로 메서드에서 호출자 인스턴스에 접근하여 값을 변경할 수 있다.
// 5. 값 타입 메서드 호출 시 값이 모두 복사된다. 인스턴스가 아닌 값 중심의 메서드를 만들 때 사용한다.
//    호출자 인스턴스에 접근할 수 없고, 복사되는 양에 따라서 성능상 문제가될 수 있다.
