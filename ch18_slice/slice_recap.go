package ch18_slice

import "fmt"

func SliceRecap() {
	// var runeSlice = []rune{'k', 'w', 'o', 'n'}
	runeSlice2 := make([]rune, 10, 9)
	for idx, val := range runeSlice2 {
		fmt.Printf("%d, %c\n", idx, val)
	}
}
