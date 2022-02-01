package main

import (
	ch3 "TuckerGolang/ch3_HelloGoWorld"
	ch4 "TuckerGolang/ch4_Variable"
	ch5 "TuckerGolang/ch5_fmt"
	ch6 "TuckerGolang/ch6_operator"
	"fmt"
)

func main() {
	chapter3()
	chapter4()
	chapter5()
	chapter6()
}

func chapter6() {
	ch6.Op_example()
}

func chapter3() {
	fmt.Println("\n\n\n\nHello golang!")
	var AddedResult = ch3.AddLiteral(10, 3)
	fmt.Println(AddedResult)
}

func chapter4() {
	var a, b int = 3, 4
	var result = ch4.VariableSum(a, b)
	fmt.Println("ch4 result: ", result)
	ch4.MemAllocAndInitialization(a, "hei")

	//var isVisible = ch4.AmIVisible
	//fmt.Println(isVisible)

	ch4.VariableType()

	//1. unsigned integer
	//uint8, uint16, uint32, uint64

	//2. signed integer
	//int8, int16, int32, int64

	//3. float
	//float32, float64

	//4. Complex number
	//complex64, complex128 - 진수, 가수는 각각 float32 범위와 같음

	//5. etc
	//byte는 uint8의 별칭.

	//	rune은 UTF-8로 문자하나 표현가능(int32의 별칭)
	//int는 32비트 CPU 컴퓨터는 int32, 64비트 CPU 컴퓨터는 int64와 동일
	//uint는 32비트 CPU 컴퓨터는 uint32, 64비트 CPU 컴퓨터는 uint64와 동일

	ch4.TypeDefaultValue()
}

func chapter5() {
	ch5.FmtPrintFunctions()
	ch5.FmtScanFunctions()
}
