package main

import "fmt"

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
}
