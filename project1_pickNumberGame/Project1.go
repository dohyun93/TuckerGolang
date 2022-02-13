package project1_pickNumberGame

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func InputIntValue() (int, error) {
	var number int
	stdin := bufio.NewReader(os.Stdin)
	_, err := fmt.Scanln(&number) // 한 줄 입력받고
	if err != nil {
		stdin.ReadString('\n') // 정수하나 입력아니라 한 줄 내 공백이나 다른 정수아닌 값 입력시, 엔터까지해서 버퍼비우기.
		fmt.Println(err)
	}
	return number, err
}

func PickNumberGame() {
	fmt.Println()
	rand.Seed(time.Now().UnixNano())
	// time.Now().UnixNano(): UTC 시간 기준 1970년 1월 1일부터 현재까지 경과한 나노초를 리턴함.
	var targetNumber = rand.Intn(100)

	var cnt int = 1
	fmt.Println("[정답]: ", targetNumber)
	for true {
		fmt.Print("숫자값을 입력하세요:")
		myNumber, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
			continue
		} else {
			if myNumber == targetNumber {
				fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: %d\n", cnt)
				break
			} else {
				if myNumber < targetNumber {
					fmt.Println("입력하신 숫자가 더 작습니다.")
				} else if myNumber > targetNumber {
					fmt.Println("입력하신 숫자가 더 큽니다.")
				}
				cnt++
			}
		}
	}
}
