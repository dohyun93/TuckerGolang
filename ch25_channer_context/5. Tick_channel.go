package ch25_channer_context

// 채널에 메시지가 있다면 빼와서 실행하고, 그렇지 않다면 1초 간격으로 다른 일을 수행해야 한다고 해보자.
// time패키지의 Tick() 함수로 원하는 시간 간격으로 신호를 보내주는 채널을 만들 수 있다.

import (
	"fmt"
	"sync"
	"time"
)

func square4(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second) // 1초 간격 시그널
	// 1. time.Tick()은 일정 시간 간격 주기로 신호를 보내주는 채널을 생성해서 반환하는 함수이다.
	// 이 함수가 반환한 채널에서 데이터를 읽어오면 일정 시간 간격으로 현재 시각을 나타내는 Time 객체를 반환한다.

	terminate := time.After(10 * time.Second) // 10초 이후 종료 시그널
	// 2. time.After()은 현재 시간 이후로 일정 시간 경과 후에 신호를 보내주는 채널을 생성해서 반환하는 함수이다.
	// 이 함수가 반환한 채널에서 데이터를 읽으면 일정 시간 경과 후에 현재 시각을 나타내는 Time 객체를 반환한다.

	for {
		select {
		// tick -> terminate -> ch 채널 순서대로 채널로부터 데이터 읽기를 시도한다.
		case <-tick:
			fmt.Println("TICKTOK (1s)")
		case <-terminate:
			fmt.Println("TERMINATE (after 10s)")
			wg.Done()
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func TickExample() {
	// time 패키지의 Tick 함수와 After 함수로 일정 간격, 최종 시점에 대한 로직을 개발할 수 있다.
	// 채널로 데이터를 받지 않을 때와 일정 시간이 지난 다음 처리를 하기 위해 사용할 수 있다.
	var wg sync.WaitGroup
	myChan := make(chan int)
	wg.Add(1)

	go square4(&wg, myChan)
	for i := 0; i < 10; i++ {
		myChan <- i * 2
	}
	wg.Wait()
}
