// package main

// import (
// 	"fmt"
// )

// type bank struct {
// 	blance int
// }

// func main() {
// 	c := make(chan int, 1)
// 	b := bank{}

// 	f1 := func(x int) {
// 		c <- x
// 		b.blance += <-ㅊ
// 	}

// 	for i := 0; i < 100; i++ {
// 		go f1(10)
// 	}
// 	fmt.Print(10)
// }
package main

import (
	"bank"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	done := make(chan bool, 10_000_000)

	// Alice
	for i := 0; i < 10; i++ {
		go func() {
			bank.Deposit(1)
			done <- true
		}()
	}

	// Wait for both transactions.
	<-done

	fmt.Printf("Balance = %d\n", bank.Balance())
	defer log.Printf("[time] Elipsed Time: %s", time.Since(start))
}
