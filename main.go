package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
	var myF float64 = 1e1
	for i := 0; i < 10; i++ {
		myF -= 0.1
	}

	fmt.Println(myF)
	var machineEpsilon = 1e-14
	if math.Abs(myF-9.0) <= machineEpsilon {
		fmt.Println("두 값은 동일합니다. 다만 이런 하드코딩방식은 옳지 않다.")
	}

}
