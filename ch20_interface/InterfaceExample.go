package ch20_interface

import (
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

type Stringer interface {
	// Fly() string -> 구현하지 않아 오류남..
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func InterfaceExample() {
	fmt.Println()
	student := Student{"도현", 30}
	var stringer Stringer
	stringer = student // interface형에 Student형 대입

	fmt.Printf("%s\n", stringer.String())
}
