package ch25_channer_context

import (
	"fmt"
	"sync"
	"time"
)

func ChannelExample() {
	// 모든 고루틴들이 종료될 때 까지 대기하는 그룹을 만든다.
	var wg sync.WaitGroup
	ch := make(chan int) // int 형 데이터를 넣고 뺄 수 있는 채널을 만들었다.

	wg.Add(1) //WaitGroup에 대기중인 고루틴의 갯수 설정
	go square(&wg, ch)
	//고루틴 생성. 채널 인스턴스를 받아서 채널에서 데이터를 받아오려고 시도하지만, 아직 채널에 데이터가 없기 때문에
	// 데이터가 들어올 때 까지 대기한다.
	ch <- 9   // 이제 데이터를 넣어줬기 때문에, 22행에서 변수 n으로 채널의 데이터를 빼올 수 있다. 26행에서 대기중인 고루틴의 개수 -1.
	wg.Wait() // 모든 고루틴이 종료될 때 까지 대기.
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch // 채널 인스턴스 ch에 담긴 데이터를 꺼내 n 변수에 담는다.

	time.Sleep(time.Second)
	fmt.Printf("Square : %d\n", n*n)
	wg.Done() // 대기중인 고루틴의 수행이 종료됨을 알려준다.
}
