package main

import (
	"TuckerGolang/Chapter4_Variable/customPkg"
	"fmt"
	"reflect"
)

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	customPkg.Declaration(3)
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var testvar float32 = 22.22 // 1.2345233e+3?
	fmt.Println(testvar)
	fmt.Println(reflect.TypeOf(testvar))

	fmt.Printf("%b", testvar)
}
