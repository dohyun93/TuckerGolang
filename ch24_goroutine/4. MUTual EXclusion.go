package ch24_goroutine

import (
	"fmt"
	"sync"
	"time"
)

// 동시성 프로그래밍을 할 때 동일 자원에 동시에 접근하는 경우 예상외 결과를 만날 수 있다.
// 이런 문제를 방지하고자, 한 자원에 대해서 동시에 접근하는 것을 방지하는 방법중에 뮤텍스(MUTual EXclusion)가 있다.

// 우리말로 직역하면 '상호 배제'다.

// 뮤텍스의 Lock() 메서드를 호출해 뮤텍스를 획득할 수 있고, Unlock() 메서드를 호출해서 뮤텍스를 반납할 수 있다.
// 그러면 대기하던 다른 고루틴이 뮤텍스를 획득해 자원에 접근할 수 있다.

var mutex sync.Mutex

type Account2 struct {
	Balance int
}

func DepositAndWithdraw(account *Account2) {
	mutex.Lock()         // 고루틴은 뮤텍스를 획득할 때까지 대기하고, 획득하면 아래 코드 실행한다.
	defer mutex.Unlock() // 무조건 함수 Deposit~ 종료 직전에 호출. (뮤텍스 반납)
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func MUTEX() {
	var wg sync.WaitGroup
	account := &Account2{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
