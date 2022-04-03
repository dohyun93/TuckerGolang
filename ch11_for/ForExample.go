package ch11_for

import (
	"bufio"
	"fmt"
	"os"
)

func ForExample() {
	// go 언어에는 유일한 반복문인 for가 있다.
	// while은 없다.

	// 초기문; 조건문; 후처리
	// 1. 초기문; 조건문;
	for i := 0; i < 10; {
		fmt.Println(i, "번째 루프입니다.!")
		i++
	}

	// 2. ; 조건문; 후처리
	var j = 0
	for ; j < 10; j++ {
		fmt.Println(j, "번쨰 루프입니다.")
	}

	// 3. ; 조건문;
	for j < 20 { // 혹은 for j<20 만 해도됨.
		j++
	}

	// 4. for true 또는 for -> 무한루프
	for true {
		j++
		if j == 100 {
			break
		}
	}
	for { // for true와 동일 -> switch문 처럼.
		j++
		if j == 200 {
			break
		}
	}

	// bufio, os를 이용한 콘솔입력 및 조건문을 통한 for문의 break, continue 연습하기.

	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요.")
			_, err2 := stdin.ReadString('\n') // 한줄 입력 \n -> 버퍼를 비우기.
			if err2 != nil {
				fmt.Println("something went wrong in go source.\n")
			}
			continue
		}
		fmt.Printf("입력한 숫자는 %d입니다.\n", number)
		if number%2 == 0 {
			fmt.Println("짝수를 입력했으므로 for문을 나갑니다.")
			break
		}
	}
	fmt.Println("for문, 콘솔입출력 연습완료!")

	// 중첩 for문을 탈출하는 방법은 크게 세 가지가 있다.
	// 1. flag를 활용한 break
	// 2. label을 활용한 break
	// 3. 함수를 활용한 중첩 감소

	// 1. flag 변수를 활용한 중첩 for문 벗어나기.
	a, b := 1, 1
	out := false
	for ; a < 10; a++ {
		for b = 1; b < 10; b++ {
			if a*b == 72 {
				out = true
				break
			}
		}
		if out {
			break
		}
	}
	fmt.Println("a * b = 72. a, b: ", a, b)

	// 2. Label을 활용한 중첩 for문 탈출하기.
	a, b = 1, 1
Label:
	for ; a < 10; a++ {
		for b = 1; b < 10; b++ {
			if a*b == 72 {
				break Label
			}
		}
	}
	fmt.Println("a * b = 72. a, b: ", a, b)

	// 3. 함수를 활용한 중첨 for문 축소. (또는 label 사용 지양하기)
	a = 1
	for ; a < 10; a++ {
		if b, found := foundB(a); found {
			fmt.Println("a * b = 72. a, b : ", a, b)
			break
		}
	}
}

func foundB(a int) (int, bool) {
	for b := 1; b < 10; b++ {
		if b*a == 72 {
			return b, true
		}
	}
	return 0, false
}
