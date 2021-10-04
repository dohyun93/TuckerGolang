package main

// func add (a int, b float) r int{
// 		return a + int(b)
// }

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	var isSame bool = math.Nextafter(a, b) == b
	return isSame
}

func main() {
	var x, y float64 = 3.0, 1.0 + 2.0
	// 아래와 같이 함수를 사용한다면 다음 단계로 수행된다.
	// 명령 포인터의 순서
	// 1. 함수를 호출한다.
	// 2. 매개변수(a, b)를 선언하고, 입력한 인숫값 x, y를 복사해서 각각 매개변수 a와 b에 넣는다.
	// 3. 로컬변수 isSame에 값을 할당, 반환한 뒤 함수가 종료되며, 사용한 지역변수 isSame이 소멸된다. 함수를 호출한 지점으로 명령 포인터가 되돌아간다.
	isSame := equal(x, y)
	fmt.Println(isSame)
}
