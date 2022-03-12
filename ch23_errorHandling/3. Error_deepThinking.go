package ch23_errorHandling

import (
	"fmt"
)

// error 타입은 인터페이스다.

// type error interface {
// 	 Error() string
// }

// 즉 어떤 타입이던 string 반환형의 Error 메서드를 구현하는 타입이면 에러로 사용할 수 있다.

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		// error 인터페이스로 사용자 정의 구조체 PasswordError를 받아준다.
		// 왜?> string 반환형의 Error()를 구현해줬으니까.
		return PasswordError{len(password), 8}
	}
	return nil
}

func AccountSystem() {
	err := RegisterAccount("Dohyun Kwon", "284569")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok { //error 타입을 PasswordError로 type assertion하여 필드에 접근가능하다.
			fmt.Printf("%v len: %d RequireLen:%d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원가입 완료.")
	}
}
