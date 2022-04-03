package ch13_struct

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32
	Score float64
}
type Student struct {
	Name     string
	Class    int
	Grade    string
	PhoneNum string
	Married  bool
}

type Person struct {
	StudentInfo Student // 구조체 내부에 구조체가 존재할 수 있다.
	// Student --> '포함된 필드'라고 하며, 이 경우 Person 타입의 변수에서 . 하나만으로 Student 구조체의 변수에 접근할 수 있다.
	// Hw.Name 가능.
	// Student
	Age    int
	Height float64
	Weight float64
}

func StructExample() {
	var dohyun Student = Student{Name: "Dohyun", Grade: "S", Class: 1, PhoneNum: "+821077934180"}
	fmt.Println(dohyun.Name)
	fmt.Println(dohyun.Married)

	var Hw Person
	Hw.StudentInfo.Name = "Hyangwon"
	Hw.StudentInfo = Student{Name: "Hyangwon", Class: 10, Grade: "S", PhoneNum: "+821031300565", Married: false}
	Hw.Age = 31
	Hw.Height = 171.3
	Hw.Weight = 68
	fmt.Println(Hw)

	// 라인 19가 활성화일 경우 포함된 필드일 경우 아래처럼 선언대입문 사용 가능.
	// Hww := Person{
	// 	    Student{},
	//  	31,
	//	    171.3,
	//	    78.3,
	// }

	// 1. 구조체 값 복사.
	student1 := Student{Name: "student1", PhoneNum: "821000000000", Class: 87, Grade: "B", Married: true}
	fmt.Println(student1)
	student2 := student1 // Go에서는 필드 각각이 아닌 구조체 전체를 한 번에 복사한다.
	fmt.Println(student2)

	student2.Married = false
	fmt.Println(student1)
	fmt.Println(student2)

	// 2. 필드 배치 순서에 따른 구조체 크기 변화
	user := User{Score: 100.0, Age: 30} // Score: float64(8bytes), Age: int32(4bytes)
	fmt.Println(unsafe.Sizeof(user))    // 12가 아니라 16 바이트.
	// '메모리 정렬'?
	// 컴퓨터가 데이터에 효과적으로 접근하고자 메모리를 일정 크기 간격으로 정렬하는 것을 말한다.
	// 실제 데이터가 연산되는 CPU 내 레지스터가 8바이트인 컴퓨터를 64비트 컴퓨터라고 하는데,
	// 이는 곧 한 번에 연산할 수 있는 크기가 8바이트라는 의미이다.

	// 따라서 데이터가 레지스터와 똑같은 크기로 정렬되어 있으면 더욱 효율적으로 데이터를 읽어올 수 있다.

	// 메모리 패딩을 통해 레지스터로 읽어오는 필드 사이에 빈 공간을 포함하더라도 메모리 정렬을 하는 이유는
	// 연산에 필요한 데이터를 한 번에 로드하여 연산을 통해 성능을 향상시키기 위함.
	memory := memory1{}
	fmt.Println(unsafe.Sizeof(memory))

	// 구조체 정리
	// 구조체는 관련된 데이터들을 묶어서 응집성을 높이고 재사용성을 증가시킬 수 있다.
	// 객체 간 결합도(의존관계)는 낮추고, 응집성을 높이는 객체지향 설계원칙인 SOLID 원칙에 부합.
	// 배열도 마찬가지로 관련된 데이터들을 묶어서 응집도를 높임.
	// 함수도 관련 코드 블록을 묶어서 응집도를 높이고 재사용성을 높임.
}

type memory1 struct {
	// 8바이트 컴퓨터라고 했을 때, 레지스터로 올리는 데이터 타입의 바이트 사용: 0. x는 데이터 패딩.
	//A int8  // 0 x x x x x x x
	//B int64 // 0 0 0 0 0 0 0 0
	//C int8  // 0 x x x x x x x
	//D int64 // 0 0 0 0 0 0 0 0
	//E int8  // 0 x x x x x x x
	//F int64 // 0 0 0 0 0 0 0 0
	// -> 결과적으로 6 * 8 = 48 바이트가 이 구조체의 메모리 할당 크기임.

	// 하지만, 이를 int8 필드들을 위로 올린다면
	A int8 // 0 0 0 x x x x x
	C int8
	E int8
	B int64 // 0 0 0 0 0 0 0 0
	D int64 // 0 0 0 0 0 0 0 0
	F int64 // 0 0 0 0 0 0 0 0
	// -> 4 * 8 = 32 바이트가 이 구조체의 메모리 할당 크기임.
}
