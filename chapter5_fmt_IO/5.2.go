package main

import "fmt"

func main() {
	var myAge int64 = 29
	var yourAge int64 = 35
	var myHeight float64 = 7.2
	fmt.Printf("%05d, %05d\n", myAge, yourAge)   // 최소 표시되는 자리수가 5인 정수를 출력. 5보다 짧은 길이의 정수면 앞에 0을 채움.
	fmt.Printf("%-05d, %-03d\n", myAge, yourAge) // 최소 표시되는 자리수가 5, 3인 정수를 출력. 이보다 짧은 길이의 정수면 뒤에 0을 채움.
	fmt.Printf("%5.3f", myHeight)                // 앞 숫자는 소수점 포함 표시되는 숫자 갯수. 뒤 숫자는 소수점 이하 표시되는 숫자갯수.

}
