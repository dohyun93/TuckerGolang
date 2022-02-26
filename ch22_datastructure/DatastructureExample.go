package ch22_datastructure

import (
	"container/list"
	"container/ring"
	"fmt"
)

func DataStructure() {
	// 22장은 자료구조로, 아래 세 개의 자료구조에 대해 배운다.
	fmt.Println()
	// 1. 리스트
	// 2. 링
	// 3. 맵
	List()
	Ring()
	Map()
}

type Element struct {
	Value interface{} // 데이터를 저장하는 필드
	Next  *Element    // 다음 요소의 주소를 저장하는 필드
	Prev  *Element    // 이전 요소의 주소를 저장하는 필드
}

// Queue //
type Queue struct {
	v *list.List // container/list의 List 구조체를 사용.
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

// Stack //
type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.l.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.l.Back()
	if back != nil {
		return s.l.Remove(back)
	}
	return nil
}

func List() {
	// list는 기본 자료구조로 여러 데이터를 보관할 수 있다.
	// 배열은 연속된 메모리에 저장하지만, 리스트는 불연속된 메모리 공간에 저장한다는 차이가 있다.

	// 1. 포인터로 연결된 요소
	// 리스트는 각 데이터를 담고있는 요소들을 포인터로 연결한 자료구조다.
	// 요소들이 서로 포인터로 연결되었다고 해서 linked list라고 부르기도 한다.

	// 2. 리스트 기본 사용법
	myList := list.New()
	element4 := myList.PushBack(4)
	element1 := myList.PushFront(1)
	myList.InsertBefore(3, element4)
	myList.InsertAfter(2, element1)
	// 1 2 3 4

	for e := myList.Front(); e != nil; e = e.Next() {
		// Next: 다음 요소의 인스턴스를 가리킴.
		fmt.Print(e.Value, " ")
	}

	for e := myList.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}

	// 3. 배열 vs 리스트
	// 배열의 경우 맨 앞에 요소를 추가할 경우, O(N)동안 한 칸 씩 뒤로 밀어내고 맨 앞에 요소값을 바꿔준다.
	// 반면 리스트는 밀어낼 필요 없이 맨 앞에 요소를 추가하고 연결만 해주면 된다.
	// 즉, 삭제나 추가는 리스트가 O(1)로 유리하다.

	// 하지만, 특정 인덱스의 요소에 접근하려면 배열은 바로 인덱스로 O(1)만에 접근이 되지만
	// 리스트의 경우 연결된 다음 인스턴스의 주소에 접근을 반복해야 하는 O(N)이 소요된다.

	// 따라서, 특정 요소에 접근할 일이 많으면 배열을, 삭제나 추가가 빈번하다면 리스트를 사용하는 것이 현명하다.

	// 데이터 지역성?
	// 데이터가 어떤 주소 근방에 밀집한 정도를 의미한다.
	// 배열은 연속된 메모리로 이뤄진 자료구조이고, 리스트는 불연속하기 때문에 배열이 리스트에 비해 데이터 지역성이 월등히 좋다.
	// 그래서 삽입과 삭제가 빈번하면 리스트를 사용하는게 좋다고 하지만, 일반적으로 요소 수가 적으면 배열이 더 효율적이다.

	// 4. 실습: 큐 구현하기
	// 리스트로 큐를 구현해보자. 큐는 선입선출의 자료구조다.
	fmt.Println()
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}

	// 5. 실습: 스택 구현하기
	// 리스트로 스택을 구현해보자. 스택은 후입선출의 자료구조다.
	fmt.Println()
	stack := NewStack()
	for i := 1; i < 5; i++ {
		stack.Push(i)
	}
	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}
}

func Ring() {
	// 링 자료구조는 환형 자료구조로, 끝이없으며 현재 위치가 중요하다.
	// container/ring
	fmt.Println()

	r := ring.New(5)

	n := r.Len() // 링 길이

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}
	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
	fmt.Println()
	// 링은 언제쓸까?
	// 실행 취소 기능 - 일정 개수의 명령 저장하고 실행 취소 (ctrl + z)
	// 고정 크기 버퍼 기능 - 데이터에 따라 버퍼가 증가되지 않고 고정된 길이로 쓸 때
	// 리플레이 기능 - 게임 등에서 최근 플레이 10초 다시 리플레이할 때.
}

func Map() {
	myMap := make(map[string]string) // 맵 생성.
	myMap["이름"] = "권도현"
	myMap["나이"] = "30"
	myMap["키"] = "187"
	myMap["몸무게"] = "90"

	fmt.Printf("%s의 나이는 %s살 이고, 키와 몸무게는 각각 %scm, %skg 이다.\n", myMap["이름"], myMap["나이"], myMap["키"], myMap["몸무게"])

	for key, value := range myMap {
		fmt.Println(key, value) // 정렬되지 않은 값이 출력된다....
	}

	// 1. map에서 요소 삭제와 없는 요소 확인
	m := make(map[int]int)

	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	m[5] = 0

	delete(m, 1)
	delete(m, 999) // 없는 키에 대해서는 아무런 동작이 일어나지 않는다.

	fmt.Println("m[1]: ", m[1]) // 이미 삭제된 키의 값은 해당 값의 기본값으로 출력.
	fmt.Println("m[5]: ", m[5]) // 0.

	// 없는 값과 실제 값이 0인 map의 요소들을 어떻게 구별할까?
	val1, ok1 := m[1]
	val5, ok5 := m[5]

	// 똑같이 값은 0이지만, ok값이 true/false로 갈린다.
	fmt.Println("m[1]의 val1, ok1: ", val1, ok1)
	fmt.Println("m[1]의 val5, ok5: ", val5, ok5)
}

// 22장 핵심 요약
// 1. 배열은 연속된 메모리 사용하고 리스트는 불연속 메모리를 사용한다.
// 2. 리스트는 요소 추가와 삭제속도가 O(1)이다. 배열은 O(N)
// 3. 링은 처음과 끝이 연결된 원형 구조다.
// 4. 맵을 사용하면 키, 값 쌍으로 되어있는 데이터를 매우 빠르게 처리할 수 있다.
