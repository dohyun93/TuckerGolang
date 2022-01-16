package main

import (
	"fmt"
)

const (
	jan int = iota
	feb
	mar
)

func main() {
	fmt.Println("test")

	var myInt int32 = 33234

	switch myInt {
	case 1:
		fmt.Println("1임")
	case 2:
	case 3:
		fmt.Println("3임")
	case 36:
		fmt.Println("find")
	}

	fmt.Println(jan, feb, mar)
}
