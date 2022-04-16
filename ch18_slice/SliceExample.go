package ch18_slice

import (
	"fmt"
	"sort"
)

// will master on 13th, Apr!

func SliceExample() {
	fmt.Println()
	// recap - len() //
	var myString = "가나다라마"
	fmt.Println("len(myString): ", len(myString)) // 크기

	var runeArray []rune = []rune(myString)

	for idx, ele := range runeArray {
		fmt.Println("[", idx, "]: ", ele)
		fmt.Printf("%c\n", ele)
	}
	fmt.Println(len(runeArray))

	// 슬라이스란?
	// 슬라이스는 고정길이의 배열과 다르게, 길이가 가변적인 '동적 배열'이다.
	// 배열을 슬라이싱하여 슬라이스를 만들 수 있다.
	// 슬라이스를 선언하는 방법, 요소에 접근/순회 하는 방법, 요소를 추가하는 방법을 알아본다.

	// 슬라이스는 선언하면 길이가 0인 슬라이스가 만들어지고, 길이를 초과해서 접근하면 런 타임 에러(패닉)가 발생한다.
	var mySlice []int

	if len(mySlice) == 0 {
		fmt.Println("길이가 0인 슬라이스 선언 완료!")
	}
	// 슬라이스 길이 밖 메모리 접근: panic 발생
	// mySlice[1] = 10
	// fmt.Println(mySlice[1])

	// [1. 슬라이스 초기화 방법]
	// 1-1. {} 를 이용하기
	var slice1 = []int{1, 2, 3}

	// 1-2. make()를 이용하기
	var slice2 = make([]int, 3) // 길이 3의 int 슬라이스 생성 (기본값 0으로 초기화)
	// copiedNum := copy(slice2, slice1) 으로 대체 가능. (슬라이스 복제)

	// [2. 슬라이스 순회]
	// 2-1. len()을 이용한 인덱스로 순회
	for i := 0; i < len(slice1); i++ {
		fmt.Print(slice1[i], ", ")
	}
	fmt.Println()

	// 2-2. range를 이용한 인덱스, 값 리턴하여 순회
	for idx, val := range slice2 {
		fmt.Println("[", idx, "]: ", val)
	}
	fmt.Println()

	// [3. 슬라이스에 요소 추가] - append() 내장함수 사용
	//slice3 := append(slice2, 4)
	//for idx, val := range slice3 {
	//	fmt.Println("[", idx, "]: ", val)
	//}
	//
	// 여러 값 추가
	slice3 := append(slice2, 4, 6, 7, 8, 9, 10)
	for idx, val := range slice3 {
		fmt.Println("[", idx, "]: ", val)
	}

	// 슬라이스 append()를 사용할 때 발생하는 예기지 못하는 문제 1.
	// <!!같은 배열을 바라보더라도, 슬라이스들 간 보유하는 데이터와 len값이 다를 수 있다.!!>
	Slice1 := make([]int, 3, 5)
	Slice2 := Slice1
	Slice2 = append(Slice2, 4)
	//Slice1정보 출력(슬라이스, len, cap):  [0 300 0 500] 4 5
	//Slice2정보 출력(슬라이스, len, cap):  [0 300 0 500] 4 5
	fmt.Println("Slice1정보 출력(슬라이스, len, cap): ", Slice1, len(Slice1), cap(Slice1))
	fmt.Println("Slice2정보 출력(슬라이스, len, cap): ", Slice2, len(Slice2), cap(Slice2))

	Slice1[1] = 300

	//Slice1정보 출력(슬라이스, len, cap):  [0 300 0] 3 5
	//Slice2정보 출력(슬라이스, len, cap):  [0 300 0 4] 4 5
	fmt.Println("Slice1정보 출력(슬라이스, len, cap): ", Slice1, len(Slice1), cap(Slice1))
	fmt.Println("Slice2정보 출력(슬라이스, len, cap): ", Slice2, len(Slice2), cap(Slice2))

	Slice1 = append(Slice1, 500)

	//Slice1정보 출력(슬라이스, len, cap):  [0 300 0 500] 4 5
	//Slice2정보 출력(슬라이스, len, cap):  [0 300 0 500] 4 5
	fmt.Println("Slice1정보 출력(슬라이스, len, cap): ", Slice1, len(Slice1), cap(Slice1))
	fmt.Println("Slice2정보 출력(슬라이스, len, cap): ", Slice2, len(Slice2), cap(Slice2))

	// 슬라이스 append()를 사용할 때 발생하는 예기지 못하는 문제 2.
	// append 함수가 호출되면 먼저 빈 공간이 충분한지 확인한다.
	// 만약 충분하지 않으면 더 큰 배열을 만들게 되는데, 일반적으로 기존 배열의 2배 cap 크기로 마련한다.
	// 그런 뒤, 기존의 배열에 있던 값들을 새로운 배열에 복사하고, 새로운 배열의 맨 뒤에 새 값을 추가한다.

	// cap은 새로운 배열의 길이 값이 되고, len은 기존 길이에 추가한 개수만큼 더한 값이되고, 포인터는 새로운 배열을 가리키는
	// 슬라이스 구조체를 반환한다.

	// 빈 공간이 없을 때 append를 하면??

	mySlice1 := []int{1, 2, 3}
	mySlice2 := append(mySlice1, 4, 5)
	fmt.Println(mySlice1, len(mySlice1), cap(mySlice1)) // [1 2 3] 3 3
	fmt.Println(mySlice2, len(mySlice2), cap(mySlice2)) // [1 2 3 4 5] 5 6
	// 위와 같이 하면, mySlice2에서 새로운 공간이 없으므로 기존 배열의 2배 크기로 마련.
	// !!!! mySlice2는 이제 mySlice1과 다른 배열을 가리키게 된다. !!!!

	// 여기서 mySlice1의 [1]값을 100으로 바꾸면 두 슬라이스의 배열의 값은 어떻게될까?
	mySlice1[1] = 1000
	fmt.Println(mySlice1, len(mySlice1), cap(mySlice1)) // [1 2 3] 3 3
	fmt.Println(mySlice2, len(mySlice2), cap(mySlice2)) // [1 2 3 4 5] 5 6
	// -> mySlice1만 바뀐다. 즉 서로다른 배열을 바라보고 있다는 의미.

	// cap을 초과하는 append를 mySlice1에 해보면 어떻게될까?
	mySlice1 = append(mySlice1, 999)
	fmt.Println(mySlice1, len(mySlice1), cap(mySlice1)) // [1 1000 3 999] 4 6
	fmt.Println(mySlice2, len(mySlice2), cap(mySlice2)) // [1 2 3 4 5] 5 6
	// mySlice2에는 아무 영향이 없으며, mySlice1의 cap은 두배로, len은 기존길이+추가한 갯수가 되었다.

	// 정리하면!
	// 어떤 슬라이스가 :=로 다른 슬라이스로 생성된 경우, 두 슬라이스의 len중 작은 인덱스에 대해 값을 바꾸면
	// 두 슬라이스 모두 변경이 일어난다. append를 하면 cap을 넘어가지 않는 한 그 슬라이스만 len이 바뀐다.
	// 이 두 슬라이스는 같은 배열을 바라본다.

	// 하지만, 만약 어떤 슬라이스가 append로 cap이 증가하면, 두 슬라이스는 더이상 같은 배열을 바라보지 않는다.

}

func Slicing() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]
	// array가 길이가 5이고, slice는 [1]부터 마지막[4]까지 배열을 사용할 수 있으므로, cap은 4가된다.

	// 배열을 슬라이싱하면 그 결과로 배열 일부를 가리키는 슬라이스를 반환한다.
	// 즉, 새로운 배열이 만들어지는게 아니라, 배열의 일부를 포인터로 가리키는 슬라이스를 만들어낸다.

	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice, len(slice), cap(slice))

	fmt.Println("array[1]의 값 100으로 변경")
	array[1] = 100
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice, len(slice), cap(slice))

	fmt.Println("slice에 77 append")
	slice = append(slice, 77, 88, 99)
	// slice = append(slice, 77, 88, 99, 111) -> 99까지만 배열에 있는 값이 바뀐다.
	// 즉, 원래 배열은 최대 길이까지만 변경되고 그 이후는 무시된다.
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice, len(slice), cap(slice))

	// 마찬가지로, 슬라이싱은 배열에 대해서뿐 아니라 슬라이스에 대해서도 할 수 있다.
	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := slice1[1:2] // [2], 1, 4
	// slice2 = slice1[:] 전체 슬라이싱
	// slice2 = array[:] 전체 슬라이싱

	// 인덱스 3개로 슬라이싱해 cap 크기 조절하기.
	// slice[시작인덱스:끝인덱스:최대인덱스]
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3:4]
	fmt.Println(slice2, len(slice2), cap(slice2)) // [2, 3], 2, 3
	// cap은 최대인덱스-시작인덱스이며, 가리키는 배열은 시작인덱스:끝인덱스 이다.
	// 3개 인자로 cap크기까지 조절할 수 있다.

}

func SliceCopyAndADDSUB() {
	// 앞서 살펴본 예제들과 같이, 같은 배열을 바라보는 슬라이스들은 버그를 유발한다.
	// 이를 방지하기 위해서 슬라이스들이 각자의 배열을 가리키도록 만들면 좋은데
	// 슬라이스의 복제, 요소추가, 요소제거를 알아보자

	// 1. 슬라이스 복제하기
	slice1 := []int{1, 2, 3, 4, 5}

	// 1-1. ... 사용하기
	slice2 := append([]int{}, slice1...)
	// slice2 := append([]int{}, slice1[0], slice1[1], ..., slice1[4])와 같다.
	// !!! 배열이나 슬라이스 뒤에 ...을 붙여주면 모든 요솟값을 넣어준 것과 같다. !!!

	// 1-2. make로 슬라이스 선언 후 range로 순회하면서 배열값 할당하기
	// 위 한 줄은 아래 네 줄로 바뀔 수 있다.
	//	slice2 := make([]int, len(slice1))
	//	for i, v := range slice1 {
	//		slice2[i] = v
	//	}

	slice1[1] = 77
	fmt.Println(slice1)
	fmt.Println(slice2)
	// slice1을 변경해도 slice2는 바뀌지 않는다.!

	// 1-3. copy() 사용하기
	// func copy(dst, src []Type) int
	// 반환 정수는 실제 복사된 요소의 개수다.
	mySlice1 := []int{1, 2, 3, 4, 5}
	mySlice2 := make([]int, 3, 10) // len: 3, cap: 10
	mySlice3 := make([]int, 10)    // len, cap: 10

	copyCnt2 := copy(mySlice2, mySlice1)
	copyCnt3 := copy(mySlice3, mySlice1)

	fmt.Println(copyCnt2, mySlice2) // len이 3이므로 3개만 복사되었다.
	fmt.Println(copyCnt3, mySlice3)

	// 2. slice 요소 삭제하기
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice3)
	deleteIdx := 2
	//for idx := deleteIdx; idx < len(slice3)-1; idx++ {
	//	slice3[idx] = slice3[idx+1]
	//}
	//slice3 = slice3[:len(slice3)-1]
	//fmt.Println(slice3)
	// 아래는 개선
	slice3 = append(slice3[:deleteIdx], slice3[deleteIdx+1:]...)
	fmt.Println(slice3)

	// 3. slice 요소 추가하기
	slice4 := []int{1, 2, 3, 4, 5}
	addIdx := 3

	slice4 = append(slice4, 0) // 1. 공간 확보
	for idx := len(slice4) - 2; idx >= addIdx; idx-- {
		slice4[idx+1] = slice4[idx]
	}
	slice4[addIdx] = 100
	// 아래는 개선
	// slice4 = append(slice4[:addIdx], append([]int{100}, slice4[addIdx:]...)...)
	fmt.Println(slice4)
	// append([]int{100}, slice4[addIdx:]...) 는 임시 슬라이드 [100, 4, 5]로,
	// 불필요한 메모리를 낭비한다.
	// 이를 방지하려면 아래처럼 할 수 있다.
	// slice4 = append(slice4, 0)
	// copy(slice4[addIdx+1:], slice4[addIdx:])
	// slice4[addIdx] = 100
}

type Student struct {
	Name string
	Age  int
}
type StudentSlice []Student // Student 구조체의 슬라이스 타입 앨리어스.

// !!sort.Interface를 사용하기 위한 아래 세 개 메서드 생성/구현!!
// 나이순으로 정렬하도록 나이 기준으로 Less 함수를 구현해주었고, 이를 통해 sort.Sort 가 나이순으로 정렬하도록 동작한다.
func (slice StudentSlice) Len() int           { return len(slice) }
func (slice StudentSlice) Less(i, j int) bool { return slice[i].Age < slice[j].Age }
func (slice StudentSlice) Swap(i, j int)      { slice[i], slice[j] = slice[j], slice[i] }

func SliceSort() {
	// sort 패키지의 Ints 함수로 정수형 슬라이스를 정렬할 수 있다.
	slice := []int{1, 2, 6, -1, -3, 0}
	sort.Ints(slice)
	fmt.Println(slice)

	// float64 타입 슬라이스는 아래처럼 정렬
	sliceFloat := []float64{1.2, -3.8, -9.6, 5.3, 7.2, -10.1}
	sort.Float64s(sliceFloat)
	fmt.Println(sliceFloat)

	// 구조체 슬라이스 정렬
	s := []Student{
		{"도현", 30},
		{"향원", 31},
		{"태란", 28},
	}
	sort.Sort(StudentSlice(s))
	// StudentSlice는 이미 sort.Interface 메서드들을 포함하고 있기 때문에, sort.Sort()인수로 사용될 수 있다.
	// StudentSlice(s)는 []Student 타입인 s를 정렬 인터페이스를 포함한 타입인 StudentSlice 타입으로 변환하는 구문이다.
	// []Student 타입은 정렬에 필요한 Less(), Len(), Swap() 메서드를 가지고 있지 않기 때문에 sort.Sort() 함수의 인수로 사용될 수 없다.
	// 그래서 []Student의 별칭 타입인 StudentSlice를 만들어서 정렬 인터페이스를 사용할 수 있도록 했다.
	fmt.Println(s)

}

// 18장 슬라이스 정리
// 1. 슬라이스는 Go에서 제공하는 동적 배열 자료구조다.
// 2. append()를 이용해 슬라이스에 값 또는 슬라이스를 추가할 수 있다.
// 3. 슬라이스가 같은 배열을 가리킬 경우, 예상치 못한 문제가 생길 수 있다.
// 4. 슬라이싱은 배열 일부를 집어내는 기능이고, 그 결과로 슬라이스가 반환된다.
// 5. append()와 슬라이싱 기능을 이용해 다양한 활용이 가능하다.
// 6. sort 패키지를 사용해서 슬라이스를 정렬할 수 있다.
// 7. 슬라이스는 실제 배열 메모리를 가리키는 포인터를 가지고 있어(Data) 배열에 비해서 메모리 사용량이나 속도에 이점이 있다.
