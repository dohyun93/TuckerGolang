package ch25_channer_context

import (
	"fmt"
)

func Size_0_channel() {
	// 1. 일반적으로 channel 생성 시 size가 0으로 생성된다.
	// size가 0으로 생성된다는 말은 채널에 데이터를 넣고나서 (택배기사에게 택배줌) 빼가지 않아 (고객에게 전달x) 대기한다는 말. (결국 멈춤)
	var myChan chan int = make(chan int)

	// 2. 9라는 정수를 정수 채널에 넣음 (택배기사에게 9 전달)
	myChan <- 9

	// 3. 꺼내가지 않아 실행 안됨.
	fmt.Println("실행 안됩니다.")
	// 이렇게 size가 0인 채널을 'unbuffered channel'이라고 한다.
}
