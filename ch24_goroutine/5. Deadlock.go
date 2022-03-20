package ch24_goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 멀티코어 컴퓨터에서는 각 코어당 한번에 하나씩 고루틴을 사용하여 성능을 향상시킬 수 있다. (컨텍스트 스위칭 없이)
// 하지만 메모리의 같은 자원을 여러 고루틴이 접근하면 프로그램이 꼬일 수 있다.
// 뮤텍스를 이용하면 동시에 고루틴 하나만 접근하도록 조정해 꼬이는 문제를 막을 수 있다. (sync.Mutex.Lock()/Unlock())
// 그러나 뮤텍스를 잘못 사용하면 성능향상도 못하고 데드락이라는 심각한 문제가 발생할 수 있다.

var waitGroup sync.WaitGroup

func Deadlock() {
	rand.Seed(time.Now().UnixNano())

	waitGroup.Add(2)
	fork := &sync.Mutex{}  // fork mutex
	spoon := &sync.Mutex{} // spoon mutex

	go diningProblem("A", fork, spoon, "포크", "수저") // A는 포크 먼저
	go diningProblem("B", spoon, fork, "수저", "포크") // B는 수저 먼저
	waitGroup.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock() // 첫 번째 도구 뮤텍스 획득 시도
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock() // 두 번째 도구 뮤텍스 획득 시도
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	waitGroup.Done()
}

// 아래는 출력 결과.
//

//GOROOT=C:\Program Files\Go #gosetup
//GOPATH=C:\Go #gosetup
//"C:\Program Files\Go\bin\go.exe" build -o C:\Users\Family\AppData\Local\Temp\GoLand\___go_build_main_go.exe C:\Go\src\TuckerGolang\main.go #gosetup
//C:\Users\Family\AppData\Local\Temp\GoLand\___go_build_main_go.exe
//B 밥을 먹으려 합니다.
//B 수저 획득
//B 포크 획득
//B 밥을 먹습니다.
//A 밥을 먹으려 합니다.
//B 밥을 먹으려 합니다.

//B 수저 획득 /////////////////////////////////
//A 포크 획득 ////////////////////////////////
//-> 이 부분에서 두 사람 모두 다른 뮤텍스를 확보하지 못했고 Go 언어에서는 데드락을 감지하여 에러를 반환함.

//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [semacquire]:
//sync.runtime_Semacquire(0x344998)
//C:/Program Files/Go/src/runtime/sema.go:56 +0x49
//sync.(*WaitGroup).Wait(0x344990)
//C:/Program Files/Go/src/sync/waitgroup.go:130 +0x6b
//TuckerGolang/ch24_goroutine.Deadlock()
//C:/Go/src/TuckerGolang/ch24_goroutine/5. Deadlock.go:21 +0x1c5
//main.chapter24(...)
//C:/Go/src/TuckerGolang/main.go:59
//main.main()
//C:/Go/src/TuckerGolang/main.go:51 +0x27
//
//goroutine 6 [semacquire]:
//sync.runtime_SemacquireMutex(0xc0000180b4, 0x0, 0x1)
//C:/Program Files/Go/src/runtime/sema.go:71 +0x4e
//sync.(*Mutex).lockSlow(0xc0000180b0)
//C:/Program Files/Go/src/sync/mutex.go:138 +0x10f
//sync.(*Mutex).Lock(...)
//C:/Program Files/Go/src/sync/mutex.go:81
//TuckerGolang/ch24_goroutine.diningProblem(0x260ead, 0x1, 0xc000018098, 0xc0000180b0, 0x2614c2, 0x6, 0x2614bc, 0x6)
//C:/Go/src/TuckerGolang/ch24_goroutine/5. Deadlock.go:29 +0x3fb
//created by TuckerGolang/ch24_goroutine.Deadlock
//C:/Go/src/TuckerGolang/ch24_goroutine/5. Deadlock.go:19 +0x147
//
//goroutine 7 [semacquire]:
//sync.runtime_SemacquireMutex(0xc00001809c, 0x0, 0x1)
//C:/Program Files/Go/src/runtime/sema.go:71 +0x4e
//sync.(*Mutex).lockSlow(0xc000018098)
//C:/Program Files/Go/src/sync/mutex.go:138 +0x10f
//sync.(*Mutex).Lock(...)
//C:/Program Files/Go/src/sync/mutex.go:81
//TuckerGolang/ch24_goroutine.diningProblem(0x260eae, 0x1, 0xc0000180b0, 0xc000018098, 0x2614bc, 0x6, 0x2614c2, 0x6)
//C:/Go/src/TuckerGolang/ch24_goroutine/5. Deadlock.go:29 +0x3fb
//created by TuckerGolang/ch24_goroutine.Deadlock
//C:/Go/src/TuckerGolang/ch24_goroutine/5. Deadlock.go:20 +0x1b2
//
//종료 코드 프로세스(으)로 완료된 프로세스
