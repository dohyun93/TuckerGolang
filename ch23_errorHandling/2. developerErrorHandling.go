package ch23_errorHandling

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		// fmt.Errorf 함수로 사용자 정의 에러 메시지를 출력할 수 있다!
		return 0, fmt.Errorf(
			"제곱근은 양수여야 합니다. f:%g", f)

		// 또는, errors 패키지의 New() 함수를 이용해서 error를 생성할 수 있다. (단 인자가 string임..)
		// return 0, errors.New("제곱근은 양수여야 합니다.")

		// [참고] %g는 지수가 크면 지수표기법(%e)으로 하고 작으면 실수 표기법(%f)으로 출력한다.
	}
	return math.Sqrt(f), nil
}

func DeveloperError() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}
