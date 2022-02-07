# 9장 공유 변수를 이용한 동시성

## 9.1 경쟁 상태
경쟁 상태(race condition)은 프로그램이 여러 고루틴의 작업 간 간섭으로 인해 올바른 결과를 반환하지 못하는 상태를 말합니다.
경쟁상태는 프로그램에 숨어 있다가 높은 부하, 특정 컴파일러, 플랫폼, 아키택처 등의 특별한 경우에만 가끔 나타나기 때문에 매우 위험하며,
재현과 진단이 어렵습니다.

```go
package bank
var balance int
func Deposit(amount int) { balance += amount }
func Balance() int { return balance }
```
위 예제는 `Deposit()`과 `Balance()` 함수를 어떤 순서로 호출하건 올바른 결과값을 반환하게 됩니다.
하지만 이를 고루틴을 사용해 비동기식으로 호출한다면 더 이상 올바른 결과를 보장할 수 없게 됩니다.

왜냐하면 `Deposit()`에 `balance` 를 증가시키는 도중에 `Balance()`를 호출하게 되면 데이터가 계산되기 전 `balance` 값을 반환할 수 있고,
이를 경쟁 상태(race condition)이라고 말합니다.

많은 개발자들이 이런 경쟁 상태를 그냥 무시하는 경우가 종종 있는데, "상호 배제 비용이 너무 크다", "이 로직은 로깅에만 쓴다.", "메시지 몇 개는 놓쳐도 상관 없다."
등으로 정당화 합니다. 이런 코드를 사용하는 컴파일러와 플랫폼에 문제가 없다면, 개발자들에게 거짓된 확신을 줄 수 있으므로 주의해야 합니다.

이러한 경쟁 상태를 피하는 방법은 go에서는 3가지가 있습니다.
1. 변수를 갱신하지 않는것.
2. 여러 고루틴에서 변수 접근을 피하는 것.
3. 상호 배제(mutual exclusion)를 사용해 고루틴이 한번에 하나씩만 접근 하도록 만든다.

### 1. 변수를 갱신하지 않는것.
1번 같은 경우 아래와 같이 변수를 갱신하는 것에서 문제가 발생합니다.
```go
var icons = make(map[string]image.Image)
func loadIcon(name string) image.Image
func Icon(name string) image.Image {
    icon, ok := icons[name]
    if !ok {
        icon = loadIcon(name)
        icons[name] = icon
    }
    return icon
}
```
위 예제에서 문제가 되는 부분은 `Icon()`함수의 if문 부분입니다.
`loadIcon(name)` 부분에서 해당 이름에 대한 Image를 가져온 후 icons[name]에 넣고 있는데,
경쟁상태가 되버리면 원치 않은 결과 값이 넣어질 수 있습니다. 

이를 방지하기 위해선 변수를 선언하고 갱신하지 않는 것인데,
```go
var icons = map[string]image.Image {
    "spades.png":   loadIcon("spades.png")
    "hearts.png":   loadIcon("hearts.png")
    "diamonds.png": loadIcon("diamonds.png")
    "clubs.png":    loadIcon("clubs.png")
}
func Icon(name string) image.Image { return icons[name] }
```
이렇게 하면 경쟁 상태가 발생해도 문제가 되지 않습니다. 

### 2. 여러 고루틴에서 변수 접근을 피하는 것.
```go
// main.go
package main

import (
	"bank"
	"fmt"
	"log"
	"time"
)

const MAX = 10_000_000

func main() {
	start := time.Now()

	done := make(chan bool)

	// Alice
	for i := 0; i < MAX; i++ {
		go func() {
			bank.Deposit(1)
			done <- true
		}()
	}

	// Wait for both transactions.
	for i := 0; i < MAX; i++ {
		if flag, success := <-done; !success && !flag {
			panic("error")
		}
	}

	fmt.Printf("Balance = %d\n", bank.Balance())
	defer log.Printf("[time] Elipsed Time: %s", time.Since(start))
}
```
```go
// bank.go
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
```
위 예제에선 `chan`을 생성해서 고루틴의 동기화를 맞추고, `init()`을 사용해 `teller()`를 하나만 생성하게끔 만들어서
경쟁 상태 문제를 해결하였습니다.

동작 방식을 살펴보면, Thread의 join형태와 유사합니다.
천만개의 고루틴을 생성 후 출력한 결과 동기화가 잘 이뤄진걸 볼 수 있습니다.
[image]

이러한 방식은 Go의 슬로건인 "메모리 공유로 통신하지 말라. 대신 통신으로 메모리를 공유해라" 입니다.
파이프라인 안의 고루틴들은 일반적으로 채널을 통해 변수 주소를 단계별로 전달해 변수를 공유합니다. 
파이프라인(채널) 단계에서 다음 단계로 변수를 전달 후 이 변수에 접근하지 않는다면, 모든 변수 접근은 순차적으로 이뤄집니다.
결과적으로 변수는 파이프라인의 한 단계에 국한되며, 다음 단계에서 반복되는 식이며, 이를 직렬 제한(serial confinement)라고 합니다.

※ 채널을 매개변수로 넣을 때, 송신과 수신을 정할 수 있다.
ex) func f(send chan <- int, reciv <- chan int)

추가로 위 예제에서 `sync.WaitGroup` 을 사용하면 채널들을 쉽게 묶어서 관리할 수 있습니다.
```go
import "sync"

func main() {
    start := time.Now()

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            bank.Deposit(1)
        }()
    }
    wg.Wait()

    fmt.Printf("Balance = %d\n", bank.Balance())
    defer log.Printf("[time] Elipsed Time: %s", time.Since(start))
}
```

### 3. 상호 배제 : sync.Mutex
채널 버퍼 용량이 1인 채널을 사용하여 최대 1개의 고루틴만 공유 변수에 접근할 수 있다.
이를 이진 세마포어라고 한다.

```go
var (
    sema = make(chan struct{}, 1) // a binary semaphore guarding balance
    balance int
)

func Deposit(amount int) {
    sema <-struct{}{} // Get token
    balance += amount
    <-sema // Release token
}
```
위 예제는 sema라는 전역변수 채널을 선언 후 채널 채널 버퍼 크기에 따라 고루틴이 대기하는 특성을 이용한 방법입니다.
다른 방법으론 `sync` 패키지의 `Mutex` 타입에서 위와 동일한 동작을 기대할 수 있습니다.
```go
import "sync"

var (
    mu sync.Mutex
    balance int
)

func Deposit(amount int) {
    mu.Lock()
    balance += amount
    mu.Unlock()
}
```
이런 식으로 Lock과 Unlock 사이 구간을 임계 영역(critical section) 이라고 합니다.
헌가지 주위해야할 점은 뮤텍스에 의해 보호되는 변수는 뮤텍스 선언 직후에 선언해야 합니다.
이 규칙을 어기는 경우에는 문서화를 해야합니다.
그리고 함수, 뮤텍스 잠금, 변수를 배열하는 것을 모니터라고 합니다.
link : https://en.wikipedia.org/wiki/Monitor_(synchronization)
link : https://medium.com/a-journey-with-go/go-monitor-pattern-9decd26fb28

Lock, Unlock을 사용할 때, 사용자의 실수로 Unlock을 빼먹는 경우가 종종 있습니다.
(코드가 길어지고 복잡해질 때)
때문에 `defer`를 사용해서 Unlock이 함수 종료 후 실행되게끔 선언하는 것이 좋습니다.
```go
func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    balance += amount
}
```

### 4. 비동기시 유의사항.
```go
// main.go
package main

import (
	"bank"
	"fmt"
	"log"
	"time"
)

const MAX = 10

func main() {
	start := time.Now()

	done := make(chan bool)

	// Alice
	for i := 0; i < MAX; i++ {
		go func() {
			bank.Deposit(1)
			done <- true
		}()
	}

	// Wait for both transactions.
	for i := 0; i < MAX; i++ {
		if flag, success := <-done; !success && !flag {
			panic("error")
		}
	}

	fmt.Printf("Balance = %d\n", bank.Balance())
	defer log.Printf("[time] Elipsed Time: %s", time.Since(start))
}
```
```go
// bank.go
package bank

var deposits = make(chan int, 10) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
```
link : https://go.dev/play/p/ajoFs80k7Vy

언듯보면 정상적인 코드로 보이지만, 위 코드는 잘못된 코드입니다. 

`deposits` 채널에 버퍼가 들어갔기 때문입니다. 

상호배제와 여러 고루틴에서 변수에 접근을 피하는것처럼 보이지만, `deposits` 채널에 버퍼를 넣어줌으로서 `main`에서 생성되는 goroutine은 비동기적으로 버퍼에 값을 넣고 종료해 버리기 때문에,

`teller()`가 연산을 다 마치기전에 함수는 끝나버리게 됩니다. 

이런경우에는 버퍼를 사용하지 않으므로써 동기화를 맞춰 문제를 해결할 수 있습니다.


### 5. Mutex 주의사항
```go
import "sync"

var mu sync.Mutex

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposits <- amount
}
func Balance() int { return <-balances }

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount)

	return true
}
```

위 코드 처럼 같은 mutex를 사용할 때, 또다시 mutex를 사용하면 dead lock이 발생 할 수 있습니다. 때문에 mutex는 중복사용하는건 좋지 않으며, 새로 만들어 주는것이 좋습니다.

