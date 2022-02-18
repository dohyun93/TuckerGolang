package ch20_interface

import (
	fedexpost "TuckerGolang/ch20_1_fedex"
	koreapost "TuckerGolang/ch20_2_koreapost"
	"fmt"
)

// 인터페이스는 우리말로 '상호작용면'으로 직역할 수 있다.
// 인터페이스를 이용하면 메서드 구현을 포함한 구체화된 객체가 아닌,
// '추상화된 객체'로 '상호작용' 할 수 있다.

// 1. 인터페이스는 메서드 집합을 포함한다.
// 2. 인터페이스 내 메서드들은 이름이 반드시 있어야 한다.
// 3. 인터페이스 내 메서드들은 매개변수나 반환형이 달라도 이름이 같아서는 안된다.
// 4. 인터페이스 내 메서드의 구현을 포함하지 않는다.

// 인터페이스도 하나의 타입이기 때문에, 키워드 type을 써준다.
type DuckInterface interface {
	Fly()
	Walk(distance int) int
}

// Go 에서는 ~er을 붙여서 인터페이스명을 만드는 것을 권장하고 있다.
// 그래서 String()메서드를 갖는 인터페이스라는 뜻으로 Stringer라고 만들었다.
type Stringer interface {
	// Fly() string -> 구현하지 않아 오류남..
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	// Sprintf는 서식에 따라 문자열을 만들어서 반환하는 함수다..
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func InterfaceExample() {
	fmt.Println()
	student := Student{"도현", 30}
	var stringer Stringer
	// Student 타입이 String() 메서드를 포함하고 있기 때문에, Stringer 인터페이스의 변수인 stringer에 대입이 가능하다!!
	stringer = student // interface형에 Student형 대입

	fmt.Printf("%s\n", stringer.String())
}

// stringer Stringer <- 인터페이스 구현 -- student Student (구조체, String 메서드)

// 인터페이스를 왜 쓰나?
// 인터페이스는 OOP에서 아주 중요한 역할을 한다. 인터페이스를 이용하면 구체화된 객체가 아닌 인터페이스 만으로도
// 메서드를 호출할 수 있기 때문에 큰 코드 수정없이 필요에따라 구체화된 객체를 바꿔서 사용할 수 있다.
// 이렇게 함으로써 변경 요청에 유연하게 대처할 수 있다.

// 추상화란?
// 추상화는 구현에 대한 구체적인 내용은 파악하지 않아도 호출하는 주체가 기능을 사용할 수 있도록 하여 유연성을 높이는 특성이다.
// 인터페이스는 호출하는 주체와 서비스를 제공하는 주체를 추상화된 관계로 서로를 이어주는 역할을 한다.

// 덕 타이핑이란?
// 인터페이스를 구현한다는 명시적 키워드 없이도 인터페이스의 메서드를 구현하는 타입의 객체가 전달되면
// 알아서 그 인터페이스를 구현하는 것으로 간주하고 별도의 다른 키워드 정의없이 인터페이스를 사용자측에서 사용할 수 있는 특징이다.

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) { // 뒤에 인터페이스로 받아주도록 구현하여 각 구조체 타입으로 넘겨주지 않는 유연함!!
	sender.Send(name)
}

func WhyUseInterfaceExample() {
	koreaPostSender := &koreapost.PostSender{}
	fedexPostSender := &fedexpost.FedexSender{}
	//
	SendBook("어린 왕자", koreaPostSender)
	SendBook("Young Prince", fedexPostSender)
}
