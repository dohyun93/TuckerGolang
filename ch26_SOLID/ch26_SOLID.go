package ch26_SOLID

import "fmt"

func Solid() {
	Liscov()
}

// < 1. 단일 책임 원칙 : 모든 객체는 하나의 책임을 진다. >
// 이점: 코드 재사용 성을 높여준다.
// 예시: 인터페이스에 대해 하나의 목적만을 갖는 객체를 구현하고, 각 객체가 맡은 독립적인 기능을 활용한다.

// < 2. 개방-폐쇄 원칙 : 확장에는 열려있지만, 변경에는 닫혀있다. >
// 이점: 상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 된다.
// 예시: 'Send' 라는 함수를 case문으로 분기처리되어있는 것을 인터페이스의 메서드로 처리하고 이를 성격에 맞게 각 객체(이메일/팩스등)가 구현하는
// 방식으로 수정한다. 이렇게 하면 추후 택배가 추가될 때 기존 객체의 구현은 건드리지 않고, 택배에 대한 객체 구현을 추가하기만 하면 된다.

// < 3. 리스코프 치환 원칙 : 상속관계에서 하위 타입에 대해서도 상위타입과 마찬가지로 동작해야 한다. >
// 이점: 예상치 못한 작동을 예방할 수 있다.
type Report interface {
	Report() string
}

type MarketingReport struct {
	title string
}

func (m *MarketingReport) Report() string {
	return m.title
}

func Liscov() {
	// 3. Liskov 치환 원칙
	var report = &MarketingReport{title: "Dohyun"} // -> panic 발생
	//var report Report // -> panic 발생 안함.
	SendReport(report)
}

func SendReport(r Report) {
	if _, ok := r.(*MarketingReport); ok { // 하위 타입이 전달될 경우 panic 발생으로 동작 안함.
		panic("Cannot send MarketingReport")
	}
	fmt.Println("Type assertion from r to *MarketingReport cannot be done.") // 상위타입이 전달되면 이 행 출력(정상동작)
}

// < 4. 인터페이스 분리 원칙 (Interface Segregation Principle) >
// 정의: "클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다."
// 이점: 인터페이스를 분리하면 불필요한 메서드들과 의존관계가 끊어져서 더 가볍게 인터페이스를 이용할 수 있다.

// 즉, 사용하지 않는 메서드들이 포함되어있는 거대한 인터페이스를 이용하려면 사용하지않는 메서드들을 불필요하게 구현해야하므로,
// 인터페이스를 적은 수의 메서드를 갖는 여러개로 분리하여 구현/사용하는 것이 좋다는 의미이다.
// 여기서 클라이언트는 객체를 의미한다.

// < 5. 의존관계 역전 원칙 (Dependency Inversion Principle) >
// 단일 책임원칙과 함께 SOLID에서 가장 중요한 원칙.
// "상위 계층이 하위 계층에 의존한느 전통적인 의존 관계를 역전시킴으로서 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다."

// 원칙 1. 상위 모듈은 하위 모듈에 의존하지 않아야 한다. 두 모듈 다 추상 모듈에 의존해야 한다.
// 원칙 2. 추상 모듈은 구체화된 모듈에 의존해서는 안된다. 구체화된 모듈은 추상 모듈에 의존해야 한다.

// 이점 1. 구체화된 모듈이 아닌 추상 모듈에 의존함으로써 '확장성' 이 증가한다.
// 이점 2. 상호 결합도가 낮아져 다른 프로그램 이식성이 증가한다.
