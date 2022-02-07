package main

import (
	"fmt"
	"sync"
)

type bank struct {
	blance int
}

func main() {
	c := make(chan int)
	b := bank{}
	var wg sync.WaitGroup

	f1 := func(x int) {
		defer wg.Done()
		c <- x
		b.blance += <-c
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(1)
	}
	wg.Wait()
	fmt.Println(b.blance)
}

// package main

// import (
// 	"bank"
// 	"fmt"
// 	"log"
// 	"time"
// )

// func main() {
// 	start := time.Now()

// 	done := make(chan bool, 10_000_000)

// 	// Alice
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			bank.Deposit(1)
// 			done <- true
// 		}()
// 	}

// 	// Wait for both transactions.
// 	for i:= 0; i < 10; i++ {
// 		if flag, success := <- done; !flag, !success {
// 			panic("error")
// 		}
// 	}

// 	fmt.Printf("Balance = %d\n", bank.Balance())
// 	defer log.Printf("[time] Elipsed Time: %s", time.Since(start))
// }
