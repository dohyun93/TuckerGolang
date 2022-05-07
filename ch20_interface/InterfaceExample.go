package ch20_interface

import (
	fedexpost "TuckerGolang/ch20_1_fedex"
	koreapost "TuckerGolang/ch20_2_koreapost"
	"fmt"
)

// 5/8 일 :
// 아침: 20~25강 복습
// 점심: go루틴활용 프로젝트2 개선

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

// 인터페이스 기능 세 가지
// 1. 포함된 인터페이스
// 2. 빈 인터페이스
// 3. 인터페이스 기본값

// 구조체에서 다른 구조체의 필드를 가질 수 있듯 인터페이스도 다른 인터페이스를 포함할 수 있다.
// 이를 포함된 인터페이스라 한다.

// [1] 포함된 인터페이스
type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader // Reader의 메서드 집합을 포함한다.
	Writer // Writer의 메서드 집합을 포함한다.
}

// [2] 빈 인터페이스
// interface{}는 메서드를 갖고있지 않은 빈 인터페이스다.
// 가지고 있어야 할 메서드가 하나도 없기 때문에 모든 타입이 빈 인터페이스로 쓰일 수 있다.
// 빈 인터페이스는 어떤 값이라도 받을 수 있는 함수/메서드/변수값을 만들 때 사용한다.

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("v is int %d\n", int(t))
	case float64:
		fmt.Println("v is float64 %f\n", float64(t))
	case string:
		fmt.Println("v is string %s\n", string(t))
	default:
		fmt.Println("Not supported type: %T:%v\n", t, t)
	}

}

// [3] 인터페이스 기본값 nil
// 인터페이스 변수의 기본값은 유효하지 않은 메모리 주소를 나타내는 nil이다.
type Strong interface {
	PrintMyValue() error
}

func CallStrong() {
	var strongVar Strong

	// strongVar은 초기화하지 않아서 기본값 nil이다.
	//그런데 nil에 접근해서 메서드 호출시도하니 런타임에러난다.
	strongVar.PrintMyValue()

}

////////////////////////////////////////////////////////////////////////////////////////////////
// [인터페이스 변환하기]
// 인터페이스 변수를 타입 변환을 통해서 구체화된 다른 타입이나 다른 인터페이스로 변환할 수 있다.
// 인터페이스 변환 사용법과 주의점을 알아보자.

// [1. 구체화된 다른 타입으로 타입 변환하기]
type Stringerr interface {
	String() string
}

type Studentt struct {
	Age int
}

func (s *Studentt) String() string {
	return fmt.Sprintf("Student Age: %d", s.Age)
}

func PrintAge(stringer Stringerr) {
	// Stringer 변수 stringer는 메서드 String()만 포함하고 있기 때문에, Age값에 접근할 수 없다.!!
	// 따라서 인터페이스를 구조체 포인터 타입으로 명시적 형 변환을 해준 뒤, 값을 조회할 수 있다.!!

	s := stringer.(*Studentt)
	// interface변수.(ConcreteType)을 하면 인터페이스 변수를 ConcreteType으로 변환하고
	// 이를 다시 선언대입문으로 변수에 할당할 수 있다.

	fmt.Printf("Age: %d\n", s.Age)

}

//func main() {
//	s := &Student{Age: 30}
//	PrintAge(s)
//}

// 인터페이스의 메서드를 구현하는 구조체 포인터만 인터페이스에서 변경이 가능하다! (구조체)

// 또한 당연히 인터페이스 변수를 변환하려는 타입이 인터페이스를 이미 포함하고 있더라도
// 다른 인터페이스를 포함하는 구조체 포인터를 인터페이스 변수로 받아서 변환하는 것은 안된다!
// A와 B가 인터페이스 포함하는 구조체 포인터라고 할 때, B를 인터페이스 변수로 받아서 A로 바꾸는건 안된다는 말.

// [2. 다른 인터페이스로 타입 변환하기]
// 인터페이스 변환을 통해 구체화된 타입 뿐 아니라 다른 인터페이스로 타입을 변환할 수 있다.
// 이때는 구조체처럼 구체화된 타입으로 변환할 때와 달리, 변경되는 인터페이스가 변경 전 인터페이스를 포함하지 않아도 된다.
// 그러나 인터페이스가 가리키는 실제 인스턴스가 변환하고자 하는 다른 인터페이스를 포함해야 한다.

type Interface1 interface {
	Read()
}

type Interface2 interface {
	Write()
}

type Dohyun struct {
}

func (d *Dohyun) Read() {

}

//func Test(inter Interface1) {
//	inter2 := inter.(Interface2)
//	// inter이 *Dohyun 가리킴.
//	// 런타임시 Write를 포함하고 있지 않은 *Dohyun을 가리키는 inter가 Interface2로 변환될 수 없음.
//}
//
//func main() {
//	inter1 := &Dohyun{} // inter1은 *Dohyun 타입.
//	Test(inter1)
//}

// [3] 타입 변환 성공 여부 반환. [중요]!!!!!!!!!!!!!!!!!!!!!!!!
// 타입변환 가능여부 확인해서 런 타임에러가 발생하지 않도록 하는 타입 변환방법
// 간단하다. 변환 리턴을 두 개로 받으면 됨.

//if inter2, ok := inter.(Interface2); ok{
//	...
//}
//
//이런식으로 하면 매우 Go스럽게 구현하는 것이다!

// 20장. Interface 핵심 요약
// 1. 인터페이스는 메서드 집합체이다.
// 2. 인터페이스에서 정의한 메서드 집합을 가진 모든 타입은 인터페이스로 쓰일 수 있다.
// 3. 덕 타이핑이란 인터페이스 구현 여부를 명시적으로 선언하는 것이 아닌 (e.g., implements), 인터페이스에서 정의한 메서드 포함여부로 판단.
// 4. 인터페이스를 사용해 추상화 계층을 만들고 관계를 통한 서비스 사용자/제공자의 상호작용을 정의한다.
// 5. 모든 타입이 빈 인터페이스 변숫값(interface{})으로 쓰일 수 있다.
// 6. 인터페이스 변환을 사용하면 인터페이스 변수를 구체화된 타입이나 다른 인터페이스로 변경할 수 있다.
