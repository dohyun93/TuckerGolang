package ch25_channer_context

// 컨텍스트는 'context'패키지에서 제공하는 기능으로, 작업을 지시할 때
// 1) 작업가능시간, 2) 작업 취소 등의 조건을 지시할 수 있는 '작업 명세서' 역할을 한다.

// 새로운 고루틴으로 작업을 시작할 때 일정 시간 동안만 작업을 지시하거나, 외부에서 작업을 취소할 때 사용한다.
// 또 작업 설정에 대한 데이터를 전달할 수도 있다.

// [1] 작업 취소가 가능한 컨텍스트
// 작업 취소 기능을 가진 컨텍스트이다.
// 이 컨텍스트를 만들어서 작업자에게 전달하면 작업을 지시한 지시자가 원할 때 작업 취소를 알릴 수 있다.

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func CancelTaskContext() {
	wg1.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	// 작업 시간을 설정하는 컨텍스트도 있다.
	// 예시: ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// 이렇게 하면 3초 뒤에 종료되게 작업지시할 수 있다.
	// 두 번째 인수의 시간이 지난 뒤, "컨텍스트의 Done() 채널에 시그널을 보내 작업 종료 요청".
	// WithTimeout 함수도 cancel을 반환하기 때문에, 원하면 작업시간 전에 언제든지 작업 취소를 할 수도 있다.

	// 1. 컨텍스트 생성
	// context.WithCancel() 함수로 취소 가능한 컨텍스트를 생성함.
	// 컨텍스트 객체와 취소 함수를 반환하며, 취소 함수를 사용해서 원할 때 취소할 수 있다.

	go PrintEverySecond(ctx)

	time.Sleep(time.Second * 5) // 2. 5초 이후에 취소 함수를 호출해 작업 취소를 알린다.
	cancel()                    // 그러면 컨텍스트의 Done() 채널에 시그널을 보내서, 작업자가 작업 취소를 할 수 있도록 알린다.

	wg1.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			// 3. 인수로 받은 ctx의 Done() 채널의 시그널이 있는지 검사. 컨텍스트가 완료될 때 Done() 채널에
			// 시그널을 넣기 때문에, 여기서 메시지를 수신하면 고루틴을 종료한다.
			wg1.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}

// 3. 특정 값을 설정한 컨텍스트
// 별도 지시사항을 주어 컨텍스트에 특정 값을 키로 값을 읽어올 수 있도록 설정할 수 있다.
// context.WithValue()

func ValueContext() {
	wg1.Add(1)
	// 1. "number"를 키로하여 값을 9로 설정한 컨텍스트를 만든다.
	ctx := context.WithValue(context.Background(), "number", 9)

	go squareWithContext(ctx)
	wg1.Wait()
}

func squareWithContext(ctx context.Context) {
	// 2. 컨텍스트 ctx의 Value() 메서드로 값을 읽어온다.
	// Value의 반환형은 interface{}이므로 type assertion으로 변환하여 사용한다.

	// 이렇게 어떤 키를 갖고 외부 지시사항을 처리할 때 양 함수간 활용할 키 이름을 서로 약속해야 한다.
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square : %d\n", n*n)
	}
	wg1.Done()
}

// [추가]
// 취소도 되고 값도 설정하는 컨텍스트는 어떻게 만들까?
// 아래처럼 할 수 있다

// ctx, cancel := context.WithCancel(context.Background())
// ctx = context.WithValue(ctx, "number", 9) -> 취소가능 컨텍스트를 인자로 감싸서 값을 설정하는 컨텍스트를 만듬.
// ctx = context.WithValue(ctx, "keyword", "Tom")
// 추가는 'go concurrency patterns'를 검색해보자.

// 25강 채널과 컨텍스트 정리
// 1. 채널은 고루틴간 메시지를 전달하는 메시지 큐이다.
// 2. 채널을 이용해서 뮤텍스 없이 동시성 프로그래밍을 할 수 있다. (역할 나누기, 영역 나누기)
// 3. 생산자 소비자 패턴은 동시성 프로그래밍에서 많이 사용된다.

// 4. 컨텍스트는 작업자에게 일을 지시할 때 사용하는 '작업 명세서'이다.
// 5. 컨텍스트를 활용해서 특정 시간동안 작업을 지시하거나 외부에서 취소할 수 있다.
// 6. 채널을 제때 닫아주지 않으면 무한히 대기하는 좀비 루틴들이 생성되어, 프로그램 성능이 저하되고 메모리 사용이 계속 증가하는
//    문제가 발생할 수 있다.
