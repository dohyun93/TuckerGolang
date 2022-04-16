package ch18_slice

import "fmt"

func SliceRecap() {
	// var runeSlice = []rune{'k', 'w', 'o', 'n'}
	//runeSlice2 := make([]rune, 10, 10)
	//for idx, val := range runeSlice2 {
	//	fmt.Printf("%d, %c\n", idx, val)
	//}
	//	fmt.Println("runeSlice2 length: len(runeSlice2)", len(runeSlice2))
	// myString := "hello"

	// len(string): 크기
	// len(slice/array): 길이

	//	fmt.Println("myString length: len(myString)", len(myString))
	// myString2 := "hello권"
	//	fmt.Println("myString2 length: len(myString2)", len(myString2))

	mySlice := make([]int, 3)
	for idx, _ := range mySlice {
		mySlice[idx] = idx
		//		fmt.Println(mySlice[idx])
	}

	// 내장함수 append로 요소 추가 가능.
	//appendedSlice := append(mySlice, 3, 4, 5)
	//fmt.Println("appendSlice: ", appendedSlice)
	//
	//slice1 := make([]int, 3, 5)
	//// slice2도 slice1과 같은 배열을 바라본다.
	//slice2 := slice1
	//fmt.Println("slice1 address: ", slice1) // [0 0 0] 3 5
	//fmt.Println("slice2 address: ", slice2) // [0 0 0] 3 5
	//
	//// slice1 값을 바꾸면 slice2도 바뀐다.
	//slice1[1] = 777
	//fmt.Println("slice1 address: ", slice1) // [0 777 0] 3 5
	//fmt.Println("slice2 address: ", slice2) // [0 777 0] 3 5
	//
	//// slice1 에 append를 하면 slice1만 바뀐다. 새로 다른 배열이 생성되는건 아님.
	//slice1 = append(slice1, 500, 500, 500)
	//fmt.Println("slice1 address: ", slice1, len(slice1), cap(slice1)) // [0 777 0 500 500 500] 6 10
	//fmt.Println("slice2 address: ", slice2, len(slice2), cap(slice2)) // [0 777 0] 3 5
	//
	//slice1[2] = 44                          // slice1의 cap이 5에서 10으로 변경됨. 두 슬라이스는 이제 더이상 같은 배열바라보지않음.
	//fmt.Println("slice1 address: ", slice1) // [0 777 44 500 500 500]  6 10
	//fmt.Println("slice2 address: ", slice2) // [0 777 0] 3 5
	//
	//slice2 = append(slice2, 666, 777, 888)  // slice2에 append해도 slice1은 변경안되는 것 알 수 있음.
	//fmt.Println("slice1 address: ", slice1) // [0 777 44 500 500 500]  6 10
	//fmt.Println("slice2 address: ", slice2) // [0 777 0 666 777 888] 6 10
	//
	//slice1 = append(slice1, 1, 2)
	//fmt.Println("slice1 address: ", slice1, len(slice1), cap(slice1)) // [0 777 44 500 500 500] 8 10
	//fmt.Println("slice2 address: ", slice2, len(slice2), cap(slice2)) // [0 777 0 666 777 888] 6 10

	// 즉 정리하면 slice는 아래같은 특징이 있다.
	// cap을 정의안하고 make로 생성 시, len과 동일한 값으로 디폴트 설정된다.

	// <len과 cap의 관계>
	// len이 cap까지 증가 가능하며 그 이상이 될 경우 len과 cap이 두배가 된다.

	// 즉 cap이 최대 치라고 하면 len은 그중 데이터가 차있는 길이다.

	// slice를 다른 slice에 := 할당하면 동일한 배열을 바라보게 된다.
	// 같은 인덱스에 데이터가 있다면 그 인덱스의 데이터 변경 시 같이 바뀐다.
	// slice 하나에 append를 하면 그 slice만 변경된다. len, cap도 필요시 그 slice만 변경된다.

	// 다른 slice의 len을 벗어나는 index의 데이터 조작은 그 index가 있는 slice만 데이터 변경이 된다.
	// 인덱스 없는 다른 slice는 아무런 변화가 없다.
	// cap이 변경되면 두 슬라이스는 더이상 같은 슬라이스가 아니다.

	//////////////////////////////
	// 만약 slice2 := slice1같은게 아니라, append로 만들어지면 어떻게 될까?
	//slice3 := make([]int, 3)
	//slice4 := append(slice3, 1, 2) // --> slice3에 append한 결과지만, slice3과 다른 배열을 바라본다!!!!
	////slice4 := make([]int, 3)
	////slice4 = append(slice4, 1, 2)
	//
	//fmt.Println("slice3: ", slice3) // [0 0 0] 3 3
	//fmt.Println("slice4: ", slice4) // [0 0 0 1 2] 5 6
	//
	//slice3[2] = 77
	//fmt.Println("slice3: ", slice3, len(slice3), cap(slice3)) // [0 0 77] 3 3
	//fmt.Println("slice4: ", slice4, len(slice4), cap(slice4)) // [0 0 0 1 2] 5 6
	//
	//slice3 = append(slice3, 7, 8, 9, 10)
	//// 한번에 너무 많이 append해서 2배의 cap으로 감당 안될 때 cap 설정 관련
	//// https://kumakichi.github.io/golang-slice-grow-strategy.html
	//fmt.Println("slice3: ", slice3, len(slice3), cap(slice3)) // slice3만 바뀜.
	//fmt.Println("slice4: ", slice4, len(slice4), cap(slice4))
	//

	///// SLICING ///////

	//array := [5]int{1, 2, 3, 4, 5}
	//slice := array[1:3]
	//
	//fmt.Println("array: ", array, len(array), cap(array))
	//fmt.Println("slice", slice, len(slice), cap(slice))
	//// 첨부한 slice원리1~2처럼 슬라이스는 가리키는 배열의 내부정보 인덱스 1부터 2까지 가리킨다. 그래서 길이는 2지만, cap은 [1]부터라 4다.
	//
	//slice[0] = 3 // 가리키는 값 변경시 array도 변경이 일어났다.!!
	//fmt.Println("array: ", array, len(array), cap(array))
	//fmt.Println("slice", slice, len(slice), cap(slice))
	//
	//slice = append(slice, 4, 5, 6) // slice의 cap을 벗어나는 순간 더이상 같은 배열을 바라보지 않는다.
	//fmt.Println("array: ", array, len(array), cap(array))
	//fmt.Println("slice", slice, len(slice), cap(slice))
	//
	//slice[0] = 777 // [0] 변경시 array가 안바뀌는걸 알 수 있다.
	//fmt.Println("array: ", array, len(array), cap(array))
	//fmt.Println("slice", slice, len(slice), cap(slice))

	/////// SLICE COPY //////////

	// slice_1과 독립적인 배열을 가리키는 slice_2를 slice_1값과 동일하게 복사하기
	slice_1 := []int{1, 2, 3, 4}
	slice_2 := append([]int{}, slice_1...)

	slice_1[2] = 777
	fmt.Println("slice_1: ", slice_1, len(slice_1), cap(slice_1))
	fmt.Println("slice_2: ", slice_2, len(slice_2), cap(slice_2))
}
