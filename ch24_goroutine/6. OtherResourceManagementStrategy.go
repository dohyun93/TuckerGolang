package ch24_goroutine

// 즉 지금 동시성프로그래밍의 문제점은 같은 메모리의 자원을 여러 고루틴들이 동시에 접근하여 사용하기 때문에 발생했다.
// 뮤텍스를 '잘못' 사용하면 동시성 프로그래밍의 장점을 사용 못하거나, 데드락이 발생했다.

// 만약 각 고루틴이 같은 자원에 접근하지 않으면 애당초 이런 문제는 발생하지 않는다.
// 각 고루틴이 서로 다른 자원에 접근하게 하려면 어떻게 해야 할까?

// 1. 영역을 나눈다
// 2. 역할을 나눈다

// 그림을 그린다고 가정해보자.
// 1. 영역을 나누는 경우, 한 도화지를 영역별로 나누어 각 고루틴들이 스케치/채색을 하게 한다.
// 이럴 경우 같은 자원에 접근하는 문제가 발생하지 않는다. -> 아래 예시로 확인해보자.

// 2. 역할을 나누는 경우, 한 도화지 전체에 대한 스케치를 A가 담당하고, 스케치가 된 다음 채색을 B가 담당한다면
// 두 개 이상의 고루틴이 같은 자원에 접근하는 문제가 발생하지 않는다. -> 25강. 채널과 컨텍스트에서 살펴본다.

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

// Job이라는 인터페이스를 구현하는 SquareJob이라는 구조체를 정의한다.
func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(time.Millisecond * 1000)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func DivideArea() {
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i} // 각기 다른 영역의 job을 할당한다.
	}

	var wg sync.WaitGroup
	wg.Add(10) // 작업개수를 설정하고,

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do() // 다른 영역에 접근하는 각 작업을 고루틴으로 실행한다. 뮤텍스 없이 동시성 프로그래밍을 문제없이 개발했다.
			wg.Done()
		}()
	}
	wg.Wait()
}

// 즉, 핵심은 고루틴 간의 간섭을 없애는 것이다.
// 영역이 아닌 역할을 나누는 방법은 25장 채널과 컨텍스트에서 배운다.

// 24장  Go routine 핵심 정리.
// 1. 고루틴은 경량 스레드로 각 코어마다 하나의 OS 스레드상에서 실행/대기 하기 때문에 컨텍스트 스위칭이 발생하지 않는다.
// 2. 멀티 코어 머신에서는 여러 고루틴을 사용하여 성능을 증가시킬 수 있다. (각 코어당 하나의 OS 스레드에서 하나의 고루틴이 실행)
// 3. 여러 고루틴이 같은 메모리 영역을 조정하면 예상치못한 문제가 발생할 수 있다.
// 4. 뮤텍스는 동시에 고루틴 하나만 자원에 접근하도록 조정한다.
// 5. 뮤텍스를 잘못 사용하면 데드락 문제가 발생할 수 있다. (포크와 나이프가 있어야 식사 가능한 상황을 상상해보자)
// 6. 작업 분할 방식 (영역을 나누는) 과 역할 분할 방식 (스케치/채색 담당 고루틴 분리)으로 뮤텍스 없이 동시성 프로그래밍을 할 수 있다.!
