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

// 이렇게 뮤텍스를 이용하면 동시에 고루틴들이 하나의 자원에 접근할 일이 발생하지 않는다.
// 하지만, 또 다른 문제가 발생할 수 있다.
// 고루틴들이 동시에 동작해야 동시 처리가 되는데 이렇게 뮤텍스로 접근제어를 해버리면 자원을 접근한 고루틴들 제외하고 다른 고루틴들은
// 모두 일을 안하고 대기하게 된다..

// 즉 성능을 향상시키려고 동시성 프로그램을 구현했지만, 성능 향상을 얻지 못하는 아이러니한 문제가 생기게 된다.

// 또, 데드락이 발생할 수 있다.
// A와 B 자원에 모두 접근해야 동작하는 함수를 고루틴으로 동시성 프로그래밍을 했다고 가정하자.
// 고루틴 '가'가 A에 접근하고 고루틴 '나'가 B에 접근했을 경우, '가'와 '나'를 포함한 어떤 고루틴도 두 자원모두를 접근할 수 없어서
// 프로그램이 멈춰버리는 데드락이 발생할 수 있다.
