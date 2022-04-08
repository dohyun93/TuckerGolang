package ch18_slice

import "fmt"

func SliceRecap() {
	var myRunes = []rune{'가', '나', '다'}
	var myString = ""

	fmt.Println("myRunes: ", myRunes)
	for _, val := range myRunes {
		fmt.Printf("%c ", val)
	}

	myString = string(myRunes) //////
	fmt.Println("\nstringed myRunes: ", myString)

	// 기본값 0으로 출력됨.
	var makeSlice = make([]int, 3)
	for _, val := range makeSlice {
		fmt.Printf("made by make keyword slice: %d\n", val)
	}

	fmt.Println(len(makeSlice), cap(makeSlice))

	makeSlice = append(makeSlice, 4)
	fmt.Println(makeSlice)

	fmt.Println(len(makeSlice), cap(makeSlice))

	Slice1 := make([]int, 3, 5) // type, len, cap
	fmt.Println(len(Slice1), cap(Slice1))
}
