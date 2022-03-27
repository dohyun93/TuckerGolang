package ch25_channer_context

// 만약 채널 하나를 바라보고 기다리고 있는데 계속 데이터가 안오면 어떡할까?

// 만약 데이터가 들어오지 않는 상황에서 다른 작업을 하거나, "여러 채널을 동시에 대기" 할 수 있다.
// 여러 채널을 동시에 대기하려면 'select문'을 사용해서 대기할 수 있다.

import (
	"fmt"
	"sync"
	"time"
)

// select문으로 여러 채널에 대한 데이터 인풋을 대기할 수 있다.
// select만 사용하면 어느 한 채널로부터 데이터를 받으면 바로 종료되기 때문에,
// 반복해서 데이터를 처리하고 싶다면 for문과 같이 사용해야 한다.

func square3(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func Select_ChannelExample() {
	var wg sync.WaitGroup
	myChan := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square3(&wg, myChan, quit)

	for i := 0; i < 10; i++ {
		myChan <- i * 2
	}
	quit <- true
	wg.Wait()
}
