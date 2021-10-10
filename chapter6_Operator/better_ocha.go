package main

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
)

func FYI() {
	fmt.Println()
	var myF float64 = 1e1
	for i := 0; i < 10; i++ {
		myF -= 0.1
	}

	fmt.Println(myF)
	var machineEpsilon = 1e-14
	if math.Abs(myF-9.0) <= machineEpsilon {
		fmt.Println("두 값은 동일합니다. 다만 이런 하드코딩방식은 옳지 않다.")
	}
	// 부동 소수점 학습! 아래는 유용 사이트
	// https://gsmesie692.tistory.com/94
}

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
	// 만약 a가 b보다 작다면 a보다 1bit큰 값을 반환. 이게 b와 같다면 1bit차이로 작은 a가 b와 같다고 간주한다.
	// 만약 a가 b보다 크다면 a보다 1bit작은 값을 반환. 이게 b와 같다면 1bit차이로 큰 a가 b와 같다고 간주한다.
	// Nextafter함수는 a와 b가 같다면 a를 반환.
	// 그 이상의 차이는 두 수가 다르다고 본다.
	// 즉, 0~1bit차이만 허용한 비교법임.
}

func main() {
	// *** 1bit라는 최소한의 오차를 허용하는 방법 *** //
	// 1. Nextafter(a, b float64) (ret float64)
	// float64 형 변수 a, b를 입력받아, a가 b보다 작다면 a를 1bit만큼 증가시키고, 그렇지 않다면 a보다 1bit작은 값 반환.

	var var1, var2, var3 float64 = 0.1, 0.2, 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", var1, var2, var1+var2)
	fmt.Printf("%0.18f == %0.18f : %v\n", var3, var1+var2, equal(var3, var1+var2))

	// *** math/big 패키지를 활용해 Float 객체로 더욱 정밀한 실수 계산하기 *** //
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	switch c.Cmp(d) {
	case 0:
		fmt.Println("c.Cmp(d) 결과는 0. c == d 정밀계산 완료")
	case -1:
		fmt.Println("c.Cmp(d) 결과는 -1. c < d 정밀 계산 완료")
	case 1:
		fmt.Println("c.Cmp(d) 결과는 1. c > d 정밀 계산 완료")
	}
	fmt.Println(reflect.TypeOf(a))

	// [출력]
	// 0.1 0.2 0.3 0.3
	// 0
	// *big.Float

	FYI()
}
