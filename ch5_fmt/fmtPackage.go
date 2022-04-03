package ch5_fmt

import (
	"bufio"
	"fmt"
	"os"
)

// 이번 5장에서는 fmt 패키지를 이용해서 터미널에 입출력을 해본다.
// fmt 패키지는 구조화된 입출력을 제공하는 패키지이다.
// fmt가 제공하는 Print(), Printf(), Println()로 표준 출력을 할 수 있고,
// Scan(), Scanf(), Scanln()로 표준 입력을 할 수 있다.

// 서식 문자.
// fmt.Printf() 내에서 %뒤에 오는 문자처럼 어떤 형태의 문자가 들어올 것이라는 표현을 할 때 사용한다.
// %d, %f, %s는 각각 정수, 실수, 문자열에 대한 서식 문자다.
// %v는 기본 서식에 맞춰 데이터 타입에 따라 기본 형태로 출력해준다.

// 5장 정리
// fmt패키지를 이용해서 표준 입/출력을 할 수 있다.
// 서식문자 (%)를 사용하면 다양한 형식으로 출력할 수 있다.
// 최소 출력너비와 소숫점 아래 숫자 개수도 지정할 수 있다.
// %v는 모든 타입의 기본 서식으로 출력한다.

// 입력받을 때 에러가 발생하면 stdin := bufio.NewReader(os.Stdin)으로 표준입력읽는 객체를 생성해서
// stdin.ReadString('\n') 으로 입력 스트림을 비워준다.

func FmtPrintFunctions() {
	// 1. fmt.Print()
	var int1 int = 100
	var int2 int = 2000
	var int3 int = 300

	fmt.Print(int1, int2, int3)

	// 2. fmt.Println()
	fmt.Println(int1, int2, int3)

	// 3. fmt.Printf()
	fmt.Printf("int1: %05d, int2: %-10d, int3: %10d\n", int1, int2, int3)
	// %10d는 왼쪽 공백포함하여 10칸으로 출력.
	// %-10d는 오른쪽 공백 포함하여 10칸으로 출력.
	// %05d는 앞에서부터 0채워서 5자리로 출력.
	// %-05d는 오른쪽에 공백 포함해서 5자리로 출력. %-5d랑 똑같이 0을 붙이지는 않음.

	var a = 324.13455
	var b = 3.14

	fmt.Printf("%08.2f\n", a) //최소너비 8칸에 소수점이하 2자리, 앞 0채우기 -> '00324.13'
	fmt.Printf("%08.2g\n", a) //최소너비 8칸에 숫자 2개만 나오고, 앞 0채우기 -> '03.2e+02'
	fmt.Printf("%8.5g\n", a)  // 최소너비 8칸에 숫자 5개만 나옴 -> '  324.13' (앞 공백 2자리)
	fmt.Printf("%f\n", b)     // %f는 디폴트로 소수점 이하 숫자 6개가 표현된다. 즉, %f는 %.6f와 같다.
}

func FmtScanFunctions() {
	// 1. fmt.Scan() -> 공백(space)과 개행을 기준으로 입력받는다.
	var a, b int
	n1, err1 := fmt.Scan(&a, &b) // space나 enter로 각 변수 입력을 구분지음.
	if err1 != nil {
		fmt.Println(n1, err1)
	} else {
		fmt.Println(n1, a, b)
	}

	// 2. fmt.Scanln()
	// '한 줄'로 입력받아서 인수로 들어온 변수 메모리 주소에 값을 채워준다. 즉 공백을 기준으로 값을 받는다.
	n2, err2 := fmt.Scanln(&a, &b)
	if err2 != nil {
		fmt.Println(n2, err2)
	} else {
		fmt.Println(n2, a, b)
	}

	// 3. fmt.Scanf()
	// 서식에 맞춰 입력받기 어렵기 때문에, 일반적으로 Scan이나 Scanln을 사용한다.
	n3, err3 := fmt.Scanf("%d, %d\n", &a, &b) // -> 정해진 형식에 맞게 입력해주어야 함. 왼쪽처럼 '정수, 정수'
	if err3 != nil {
		fmt.Println(n3, err3)
	} else {
		fmt.Println(a, b)
	}

	// 추가
	// 입력 스트림을 지우기
	stdin := bufio.NewReader(os.Stdin)
	//stdin := bufio.NewReader(os.Stdin)
	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println("error: ", err)
		streamString, err2 := stdin.ReadString('\n')
		if err2 != nil {
			fmt.Println("[ERROR] cleanup input buffer: ", err2)
		} else {
			fmt.Println(streamString) // 입력 버퍼에 남은 데이터들.
			fmt.Println("[SUCCESS] cleanup input buffer.")
		}

	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(n, a, b)
	}

	// 입력 Hello 4\n -> H를 보고 79라인에서 err가 nil이 아니라 들어감. stdin이 \n가 나올때까지 읽는다. -> 표준입력 스트림 비워짐.
	//error:  expected integer
	//3 4 -> (enter) 4 3 이 입력되고 Scanln에 맞게 들어감.
	//2 3 4

	// stdin = bufio.NewReader(os.Stdin)
	// remainingBuffer, err := stdin.ReadString('\n')
}
