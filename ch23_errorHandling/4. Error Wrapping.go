package ch23_errorHandling

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 에러를 감싸서 새로운 에러를 만들어야 할 경우도 있다.
// 예) 파일에서 텍스트를 읽어서 특정 타입의 데이터로 변환하는 경우, 1) 파일 읽기 에러 필요하고 2) 타입 변환 에러가 몇 줄의 몇 칸에서 발생한지
// 알면 더 유용할 것이다. 이럴 때 파일 읽기에서 발생한 에러를 감싸고 그 바깥에 줄과 칸 정보를 넣으면 된다.

// https://wookiist.dev/102 -> golang의 문자열 읽기 참고.

func MultipleFromString(str string) (int, error) {
	// MultipleFromString는 문자열에서 두 단어를 읽어서 숫자로 변환한 뒤 곱한 결과를 반환하는 함수이다.

	// 1. 한 단어씩 읽는 bufio 패키지의 Scanner를 만들어준다.
	// NewScanner는 io.Reader 인터페이스를 인수로 받기 때문에, 문자열을 전달하기 위해 strings.NewReader() 함수를 사용했다.
	// 문자열을 한 줄씩 또는 한 단어씩 끊어 읽고자 할 때 주로 사용되는 구문이니 기억해두자.
	scanner := bufio.NewScanner(strings.NewReader(str))

	// 2. Scanner 객체의 Split 메서드를 호출 해 어떻게 끊어읽을지 알려준다.
	// bufio.ScanWords 를 사용하면 한 단어씩 끊어읽게 되고, bufio.ScanLines 를 사용하면 한 줄씩 끊어읽게 된다.
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		// Wrapping errors with %w
		// 6. readNextInt 함수에서 발생한 에러를 다시 감싸서 pos 정보를 에러에 추가했다.
		return 0, fmt.Errorf("step 1. Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("step 2. Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { // 3. Scan 함수로 첫 번째 단어를 읽어온다.
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	// scanner로 읽은 첫 단어를 저장한다.
	word := scanner.Text()

	// 4. strconv.Atoi 함수로 문자열을 int로 변환한다. Itoa 는 반대로 숫자를 문자열로 변환한다.
	// Atoi는 숫자가 아닌 문자가 섞여있는 경우에는 NumError 타입의 에러를 반환한다.
	number, err := strconv.Atoi(word)

	if err != nil {
		// 5. fmt.Errorf()함수 안에 %w 서식문자를 통해 NumError 에러(변수 err)를 감싼다.
		// --> 따라서, 다른 정보와 감싸진 에러까지 에러 하나로 반환할 수 있게 된다!!
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s, err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println("Atoi 함수 에러 발생: ", err)
		var numError *strconv.NumError
		// 7. 감싸진 에러가 NumError 인지 확인하기위해 As 함수로 꺼낼 수 있다.
		// errors.As 는 strconv.Atoi() 함수에서 발생한 에러 err가 *strconv.NumError 타입이기 때문에 true를 반환하고,
		// 변수 numError에 *strconv.NumError 타입값인 err의 값을 넣어준다.
		if errors.As(err, &numError) { // 감싸진 에러가 NumError인지 확인
			fmt.Println("NumberError:", numError)
		}
		// errors 패키지는 As() 함수 외 단순히 에러 객체 타입인지만 확인하는 Is() 함수도 있다.
	}

}

func ErrorWrappingExample() {
	readEq("123 3")
	readEq("123 abc")
}
