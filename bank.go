package bank

import "sync"

var deposits = make(chan int, 10_000_000) // send amount to deposit
var balances = make(chan int)             // receive balance
var mu sync.Mutex

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposits <- amount
}
func Balance() int { return <-balances }

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
