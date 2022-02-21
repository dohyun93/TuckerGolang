package ch22_datastructure

func DataStructure() {
	// 22장은 자료구조로, 아래 세 개의 자료구조에 대해 배운다.

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

func List() {
	// list는 기본 자료구조로 여러 데이터를 보관할 수 있다.
	// 배열은 연속된 메모리에 저장하지만, 리스트는 불연속된 메모리 공간에 저장한다는 차이가 있다.

	// 1. 포인터로 연결된 요소
	// 리스트는 각 데이터를 담고있는 요소들을 포인터로 연결한 자료구조다.
	// 요소들이 서로 포인터로 연결되었다고 해서 linked list라고 부르기도 한다.

	// 2.

	// 3.

	// 4.

	// 5.
}

func Ring() {

}

func Map() {

}
