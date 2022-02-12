package ch12_array

import (
	"fmt"
)

func ArrayExample() {
	// 배열(array)은 같은 타입의 여러 데이터들로 이루어진 타입이다.
	// 배열을 이루는 각 데이터를 '요소'라고 칭하고, 위치를 '인덱스'라고 한다.

	var myArray [10]int64 = [10]int64{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(myArray)
	// 지정 안한 나머지 인덱스들의 요소는 0으로 초기화된다.
	// [0 0 0 0 0]

	var myArray2 [5]float64
	fmt.Println(myArray2)
	// [0 0 0 0 0]

	myArray2[2] = 7.7
	fmt.Println(myArray2)
	// [0 0 7.7 0 0]

	// 1. 배열 사용법
	// 1-1. 선언 / 초기화
	var strArray [3]string = [3]string{"my", "name", "dohyun"} // 가장 정석적인 선언/초기화
	strArray2 := [4]string{"1", "2", "3", "4"}                 // 4개 string 요소 배열을 선언대입문으로 배열변수 생성
	floatArray1 := [3]float64{1: 3.7, 2: 2.5}                  // 인덱스: 요소값으로 초기화 할 수 있음.
	intArray := [...]int{1, 2, 3}                              // 배열요소 개수 -> 초기화 하는 요소의 개수로 선언가능.

	fmt.Println(strArray)
	fmt.Println(strArray2)
	fmt.Println(floatArray1)
	fmt.Println(intArray)

	// 2. range 순회
	// for 반복문에서 range 키워드를 이용하면 배열 요소를 순회할 수 있다.
	for idx, val := range intArray {
		// 'range 배열' 로 index, 요소값 반환.
		fmt.Println("index ", idx, ": ", val)
	}

	// !!! 배열은 연속된 메모리다. !!!
	// 컴퓨터는 인덱스와 타입 크기를 사용해서 메모리 주소를 찾는다.
	// [0][1][2][3][4][5]
	// A번지가 [0]배열요소의 주소(=배열이름의 주소)라면, [2] 배열요소의 주소는 [0]주소 + (2-0)*배열타입이다. (연속된 메모리이기때문)

	// 3. 배열 복사
	// 대입 연산자를 사용하면 배열 대 배열 복사를 할 수 있다.

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{100, 200, 300, 400, 500}
	fmt.Println("a 배열 요소 출력")
	for idx, val := range a {
		fmt.Println("a[", idx, "] : ", val)
	}
	fmt.Println()

	fmt.Println("b 배열 요소 출력")
	for idx, val := range b {
		fmt.Println("b[", idx, "] : ", val)
	}

	b = a // 배열 <- 배열 할당.
	// (a배열의 주소에 있는 값들이 b배열의 주소에 있는 값들로 복사된다.)
	// 얕은 복사
	a[4] = 10000
	fmt.Println(a[4], b[4]) // 10000, 5

	////
	fmt.Println("배열 b로 배열 a할당 (b = a) 한 결과:")

	fmt.Println("b 배열 요소 출력")
	for idx, val := range b {
		fmt.Println("b[", idx, "] : ", val)
	}
	// 할당받는 배열이 할당해주는 배열의 길이와 다를경우 런타임에러 발생.
	////

	// 4. 다중 배열
	var myMap [3][4]int = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}, // 마지막 }가 같은줄에 있지 않은 경우 반드시 ','로 끝내줘야 한다.!!
	}
	// 2차원 중첩배열 출력해보기.
	for _, arr := range myMap {
		for _, v := range arr {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
	// 핵심 요약
	// 1, 배열은 값을 여러개 저장하는 연속된 메모리 공간이다.

	// 2, 배열 변수는
	// var 배열이름 [3][2]int = [3][2]int{값초기화} 처럼 하거나,
	// var 배열이름 [3][2]int 로 선언만 하거나,
	// 배열이름 := [3][2]int{값초기화} 처럼 할 수 있다.

	// 3, len()으로 배열 길이를 알 수 있다.
	// 4, range를 이용하면 배열 순회를 하여 인덱스, 요소값을 리턴해 사용할 수 있다.

	fmt.Println(len(myMap))    // 3
	fmt.Println(len(myMap[0])) // 4
}
