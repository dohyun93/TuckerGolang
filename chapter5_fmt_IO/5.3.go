package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//	1. fmt.Scanf("%d %d", &a, &b) -> 포맷에 맞게 두 정수를 띄워서 입력받아야 한다.
	var a, b int
	n1, err1 := fmt.Scanf("%d %d", &a, &b)
	if err1 != nil {
		fmt.Println(n1, err1)
	} else {
		fmt.Println(n1, a, b)
	}

	// 2. fmt.Scan(&c, &d) -> 정수 두 개를 !!!'개행' 또는 '공백'!!!으로 입력받음. 더 플렉서블함. (추천)
	var c, d int
	n2, err2 := fmt.Scan(&c, &d)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(n2, c, d)
	}

	// 3. fmt.Scanln(&e, &f) -> 정수 두 개를 !!!'공백'!!!으로 입력받음.
	var e, f int
	n3, err3 := fmt.Scanln(&e, &f)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(n3, e, f)
	}

	// 4. 표준입력스트림은 FIFO 자료구조이다.
	// var a, b int 일 떄,
	// fmt.Scanln(&a, &b) 에서 Hello 4 를 입력한다면 \n4 olleH 가 입력스트림으로 차례대로 입력되는데, 정수형에 문자가 가는순간
	// 에러가 반환되며 입력스트림에는 \n4 olle 가 남게된다. 이후에 Scan 류의 함수를 사용하면 그대로 남아있던 입력스트림의 데이터가
	// 활용되므로, 입력스트림을 비워주어야 하는 작업이 필요하다. 아래는 그 방법을 나타낸 것.

	// 표준입력스트림을 인자로 받는 NewReader 함수.
	// bufio는 인자의 입력스트림으로부터 한 줄을 읽는 Reader 객체를 제공한다.
	stdin := bufio.NewReader(os.Stdin)

	var myInt int
	var myString string

	n4, err4 := fmt.Scanln(&myInt, &myString)
	if err4 != nil {
		fmt.Println("입력 중 오류 발생! ", err4)
		fmt.Println("입력스트림을 비워줍니다.")
		stdin.ReadString('\n') // 입력스트림을 비워준다.
	} else {
		fmt.Println(n4, "개 입력됨. 각 데이터는: ", myInt, myString)
	}

	// 그 이후 또 입력되는 케이스 가정
	n5, err5 := fmt.Scan(&myInt, &myString)
	if err5 != nil {
		fmt.Println("이후 입력 중 오류발생! ", err5)
	} else {
		fmt.Println(n5, "개 입력 됨. 각 데이터는: ", myInt, myString)
	}

	//<입력1>
	//hello 3
	//<출력1>
	//입력 중 오류 발생!  expected integer
	//입력스트림을 비워줍니다.

	//<입력2>
	//2 hello
	//<출력2>
	//2 개 입력 됨. 각 데이터는:  2 hello
}
