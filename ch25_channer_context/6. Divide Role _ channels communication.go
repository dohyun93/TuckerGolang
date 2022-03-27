package ch25_channer_context

// 고루틴에서 뮤텍스를 사용하지 않고 동시성 처리를 하는 방법은
// 1. 25장의 '6. OtherResourceManagementStrategy'처럼 서로 다른 영역을 처리하거나,
// 2. 아래에서 살펴볼 역할을 나누는 방벙이 있다.

// 자동차를 만들 때 1) 차체 생산 2) 바퀴 설치 3) 도색
// 의 3개 단계가 있다고 가정해보자.

// 각 작업자는 역할을 나누어 수행할 때, 최초 3초가 지나고 난 이후부터는 매 1초마다 차량이 생산된다.
// 이렇게 컨베이어벨트 시스템 처럼 역할을 나누어 동시성 처리가 되는 것을 채널을 이용해서 구현해보자.

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var StartTime = time.Now()

func DivideRoleWithChannels() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh) // 10초 이후 다음 단계인 tire 채널을 닫아주고 루틴을 종료한다.
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car // 바퀴 설치 후 paintCh 채널에 넣어준다.
	}
	// 만약 tireCh가 닫히면 루틴을 종료하고, paintCh 채널을 닫아준다.
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {
	// 마찬가지로 paintCh로 들어온 데이터들을 꺼내 도색해주고, 소요 시간과 생산된 차량정보를 출력한다.
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Blue"
		duration := time.Now().Sub(StartTime)
		fmt.Printf("%.2f Complete car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	// 만약 paintCh가 닫히면 루틴을 종료한다.
	wg.Done()
}

// 이처럼 한쪽에서 데이터를 생성해서 넣어주면 다른 쪽에서 생성된 데이터를 빼서 사용하는 방식을 '생산자 소비자 패턴' 이라고 한다.
// 여기서 MakeBody()루틴이 생산자, InstallTire 루틴은 소비자가 되며,
// InstallTire 루틴은 PaintCar 루틴에 대해서는 생산자가 된다.
