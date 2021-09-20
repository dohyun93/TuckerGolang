package main

import "fmt"

func main() {
	var myAge int64 = 29
	var yourAge int64 = 35
	var myHeight float64 = 3.1415
	var myFloat float64 = 234.12349
	// 1. 정수와 실수의 출력
	// %d, %f -> 정수와 실수 출력. fmt.Printf 함수로 사용. %f == %.6f
	// %2d -> 최소 너비 2자리 정수
	// %03d -> 최소 너비 3자리 정수. 그 미만일 경우 앞에 0을 붙여 3자리로 만듬.
	// %-3d -> 최소 너비 3자리. 그 미만일 경우 숫자를 앞에 두고 뒤에는 빈 자리로 출력.

	// %05.2f -> 최소 너비 5자리 (소수점 포함). 소수점 이하는 2자리 표현, 소수점 이하가 그 미만일 경우 뒤에 0을 채움.
	// %2.6g -> 최소 너비 2자리 (소수점, e, +, 앞 0 포함). 총 숫자 6자리.

	fmt.Printf("%05d, %05d\n", myAge, yourAge) // 최소 표시되는 자리수가 5인 정수를 출력. 5보다 짧은 길이의 정수면 앞에 0을 채움.
	fmt.Printf("%-5d, %-3d\n", myAge, yourAge) // 최소 표시되는 자리수가 5, 3인 정수를 출력. 0옵션을 줘도 적용되지 않음.
	fmt.Printf("%05.2f\n", myHeight)           // 최소 너비 5. 소수점 이하 3자리.
	fmt.Printf("%7.2g\n", myFloat)             // 최소 너비 7. 총 숫자 2자리.
}
