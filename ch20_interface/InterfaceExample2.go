package ch20_interface

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

}

// [3] 인터페이스 기본값
