package main

import (
	"TuckerGolang/Chapter4_Variable/customPkg"
	"fmt"
)

func main() {
	var day int = 3
	var month string = "Sept"

	fmt.Println(month, day, "th")

	day = 8
	month = "Nov"
	fmt.Println(month, day, "th")

	customPkg.Declaration()
}
