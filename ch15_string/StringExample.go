package ch15_string

import (
	"fmt"
)

func StringExample() {
	fmt.Println()
	// 문자열은 말 그대로 문자 집합이다.
	// 큰 따옴표(")와 백쿼트(`)를 사용해서 문자열을 출력해보자

	var string1 string = "Hi\t My name is dohyun\n"
	var string2 string = `Hi\t My name is dohyun\n`
	fmt.Println("string1: ", string1) // string1:  Hi     My name is dohyun
	fmt.Println("string2: ", string2) // string2:  Hi\t My name is dohyun\n

	// 백쿼트를 이용한 문자열 출력은 특수문자(\t나 \n같은)를 그대로 출력해버린다.
	// 백쿼트는 여러줄의 문자열도 묶을 수 있다.
	var string3 string = "안녕\n하세\n요."
	var string4 string = `안녕
하세
요.`
	fmt.Println("string3: ", string3)
	fmt.Println("string4: ", string4) // string3와 동일한 결과.

	// 1. UTF-8 문자코드
	// Go는 UTF-8 문자코드를 표준 문자코드로 사용한다.
	// UTF-8은 다국어 문자를 지원하고 문자열 크기를 절약할 목적으로 채택되었다.
	// UTF-8은 자주사용되는 영문자, 숫자, 일부 특수문자를 1바이트로, 그 외 문자들은 2~3바이트로 표현.

	// rune 타입으로 한 문자 담기
	// 문자 하나를 표현하는 데 rune 타입을 사용한다.

	var character rune = '권'
	fmt.Printf("%T\n", character) // type 출력
	fmt.Println(character)        // character값 출력 (int32)
	fmt.Printf("%c\n", character) // 문자 출력

	// 2. len()으로 문자열 '크기' 알아내기
	// 참고로, len은 12장에서 배열을 공부할 때, 배열 길이를 리턴하는 기능도 했었다.
	var korString string = "가나다라마"
	var enString string = "abcde"
	fmt.Println("len(korString): ", len(korString)) // 15bytes
	fmt.Println("len(enString): ", len(enString))   // 5bytes
	// UTF-8에서 한글은 글자당 3바이트를 차지하기 때문에, 3x5 = 15바이트가 크기이고,
	// 영문자는 글자당 1바이트이므로, 1x5 = 5바이트의 크기를 확인할 수 있다.

	// 3. []rune 타입 변환으로 글자 수 알아내기
	// string 타입, rune 슬라이스 타입인 []rune 타입은 상호 타입 변환이 가능하다.
	// [@참고@] '슬라이스'는 18장에서 자세히 다루며, 여기서는 단순히 가변길이의 배열이라고만 알고있자.
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	fmt.Println("runes: ", runes)
	fmt.Println("stringified runes: ", string(runes))
	// string 타입과 []rune 타입은 모두 문자들의 집합을 나타내기 때문에 상호 타입 변환이 가능한 것이다.

	// string 타입에 rune 슬라이스타입으로 변환하면 문자열의 길이를 알 수 있다.
	str := "롯데월드"
	runes_ := []rune(str)
	fmt.Println("문자열을 rune 슬라이스로 변환하여 각 문자의 rune타입 값들 []rune(str): ", runes_)
	fmt.Println("정수화 된 문자들의 rune타입 슬라이스 길이 len([]rune(str)): ", len(runes_))
	fmt.Println("원래 문자열의 크기인 len(str): ", len(str))

	// [@참고@] string <-> []byte 타입 변환 가능하다. 20장 참고.
	// 4. 문자열 순회
	// 4-1. 인덱스를 활용한 바이트 단위 순회
	// 4-2. []rune 타입으로 변환 후 한 글자씩 순회

	myString := "Hello 월드!"
	for _, val := range myString {
		fmt.Printf("타입: %T, 값: %d, 문자: %c", val, val, val)
		fmt.Println()
	}
	fmt.Println()
	runifiedString := []rune(myString)
	for i := 0; i < len(runifiedString); i++ {
		fmt.Printf("타입: %T, 값: %d, 문자: %c", runifiedString[i], runifiedString[i], runifiedString[i])
		fmt.Println()
	}
	// []rune으로 string을 변환시키는 과정에서 별도의 배열을 할당하므로, 추가적인 메모리 소비가 있다.
	// 따라서 위에서 range로 순회하여 출력시키는 방법이 더욱 낫다.

	// 5. 문자열 합치기
	// +로 붙일 수 있다.

	// 6. 문자열 비교하기
	// ==와 !=로 비교 가능.
	// 동일한 문자열이면 ==결과가 true가 되고, !=연산은 false가 된다.
	Str1 := "가나다"
	Str2 := "마바사"
	Str3 := "가나다"
	if Str1 == Str3 {
		fmt.Println("Str1, Str3는 같다.")
	}
	if Str2 != Str3 {
		fmt.Println("Str2, Str3가 다르다.")
	}
	if Str1 == Str2 {
		fmt.Println("Str1, Str2는 같다.")
	}

	// 7. 문자열 대소 비교하기. >, <, <=, >=
	// 첫 글자부터 값을 비교하는데, 각 문자열의 문자의 유니코드값을 기준으로 대소비교를 한다.
	yourString1 := "BBA"
	yourString2 := "BB"
	yourString3 := "BAA"
	fmt.Printf("yourString1 (%s) > yourString2 (%s) : %v\n", yourString1, yourString2, yourString1 > yourString2)
	// 없는것과 비교하면 더 크다고 간주한다. (BBA > BB)

	fmt.Printf("yourString2 (%s) > yourString3 (%s) : %v\n", yourString2, yourString3, yourString2 > yourString3)

	// 중요! - 내장함수 len()
	// len(문자열): 문자열의 크기
	// len(슬라이스 또는 배열): 슬라이스 또는 배열의 '길이'
}
