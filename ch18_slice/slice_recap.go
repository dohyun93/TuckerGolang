package ch18_slice

import "fmt"

func SliceRecap() {
	// var runeSlice = []rune{'k', 'w', 'o', 'n'}
	runeSlice2 := make([]rune, 10, 10)
	for idx, val := range runeSlice2 {
		fmt.Printf("%d, %c\n", idx, val)
	}
	fmt.Println("runeSlice2 length: len(runeSlice2)", len(runeSlice2))
	myString := "hello"

	// len(string): 크기
	// len(slice/array): 길이

	fmt.Println("myString length: len(myString)", len(myString))
	myString2 := "hello권"
	fmt.Println("myString2 length: len(myString2)", len(myString2))

	mySlice := make([]int, 3)
	for idx, _ := range mySlice {
		mySlice[idx] = idx
		fmt.Println(mySlice[idx])
	}
}
