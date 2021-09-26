package main

import (
	"fmt"
	"math/big"
	"reflect"
)

func main() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d)) // -1은 c < d, 1은  c > d.

	fmt.Println(reflect.TypeOf(a))

	// [출력]
	// 0.1 0.2 0.3 0.3
	// 0
	// *big.Float
}
