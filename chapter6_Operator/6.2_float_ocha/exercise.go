package main

import "fmt"

const epsilon = 1e-6

func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func main() {
	var a float64 = .1
	var b float64 = .2
	var c float64 = .3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f %v\n", c, a+b, equal(a+b, c))

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
}
