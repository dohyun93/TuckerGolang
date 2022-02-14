package main

import (
	ch10 "TuckerGolang/ch10_switch"
	ch11 "TuckerGolang/ch11_for"
	ch12 "TuckerGolang/ch12_array"
	ch13 "TuckerGolang/ch13_struct"
	ch14 "TuckerGolang/ch14_pointer"
	ch15 "TuckerGolang/ch15_string"
	ch16 "TuckerGolang/ch16_package"
	ch18 "TuckerGolang/ch18_slice"
	ch19 "TuckerGolang/ch19_method"
	ch3 "TuckerGolang/ch3_HelloGoWorld"
	ch4 "TuckerGolang/ch4_Variable"
	ch5 "TuckerGolang/ch5_fmt"
	ch6 "TuckerGolang/ch6_operator"
	ch7 "TuckerGolang/ch7_function"
	ch8 "TuckerGolang/ch8_const"
	ch9 "TuckerGolang/ch9_if"
	proj1 "TuckerGolang/project1_pickNumberGame"
	"fmt"
)

func main() {
	//chapter3()
	//chapter4()
	//chapter5()
	//chapter6()
	//chapter7()
	//chapter8()
	//chapter9()
	//chapter10()
	//chapter11()
	//chapter12()
	//chapter13()
	//chapter14()
	//chapter15()
	//chapter16()
	//project1()
	//chapter18()
	chapter19()
}

func chapter19() {
	// ch19.MethodExample()
	ch19.MethodVSFunction()
}

func chapter18() {
	//ch18.SliceExample()
	//ch18.Slicing()
	//ch18.SliceCopyAndADDSUB()
	ch18.SliceSort()
}
func project1() {
	proj1.PickNumberGame()
}

func chapter16() {
	ch16.PackageExample()
}

func chapter15() {
	ch15.StringExample()
}

func chapter14() {
	ch14.PointerExample()
}

func chapter13() {
	ch13.StructExample()
}
func chapter12() {
	ch12.ArrayExample()
}
func chapter11() {
	ch11.ForExample()
}

func chapter10() {
	ch10.SwitchExample()
}

func chapter9() {
	ch9.IfExample()
}

func chapter8() {
	ch8.ConstExample()
}

func chapter7() {
	ch7.FunctionExample()
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
