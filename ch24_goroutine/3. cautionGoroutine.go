package ch24_goroutine

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

// 동시성 프로그래밍을 할 때 주의해야 할 점은, 같은 자원에 동시에 접근할 경우 예상치 못한 결과를 만날 수 있다는 것이다.
// 아래 예시는 통장 잔고라는 공통된 자원에 접근해 입금과 출금을 반복하는 고루틴들이 존재할 때 발생할 수 있는 문제에 대한 예시다.
func CautionCuncurrentGoroutines() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	account := &Account{0}
	for i := 0; i < 10; i++ {
		go InandOut(account)
	}
	wg.Wait()
}

func InandOut(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should be positive value: %d", account.Balance))
	}
	account.Balance += 1000 // 1) 통장잔고를 읽어와 1000을 더하는 것과 2) 다시 잔고에 저장하는 스텝으로 나뉠 수 있다.
	// 만약 두 고루틴이 존재한다고 생각해보자.
	// 고루틴1이 값을 읽어서 1000을 더한 뒤, 아직 메모리에 쓰지 않았을 때 고루틴2가 값을 읽으면 0으로 읽을 것이다.
	// 그럼 고루틴1이 1000을 쓸 것이고, 고루틴2도 0을 읽은것에 1000을 더해 저장하니 1000으로 저장된다.
	// 즉 한 번만 더해지게 되고, 뺄 때 두 번 빼질 수 있다.
	// 그렇게 되면 다음 고루틴이 실행될 때 panic조건에 걸려 예상치 못한 결과로 이어질 수 있다.

	// 따라서, 동시성 프로그래밍을 할 때는 이처럼 자원에 접근하는 것에 대한 제한이 필요하다.
	// 4번째 파일에서 뮤텍스(MUTual EXclusion)를 알아보자.
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
	wg.Done()
}
