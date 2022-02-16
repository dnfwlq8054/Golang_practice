// main.go
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const MAX = 1000

type Memo struct {
	f     Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)

// func fu(key string) {
// 	return interface{}, error
// }

type result struct {
	value interface{}
	err   error
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	var once sync.Once
	ch := make(chan int, 1000)
	var count int

	for i := 0; i < MAX; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("hi")
			})
			ch <- 1
		}()
	}
	wg.Wait()
	close(ch)

	for i := range ch {
		count += i
	}

	var c = &Memo{cache: make(map[string]result)}
	fmt.Printf("count = %d\n", count)
	fmt.Println(c)
	defer log.Printf("[time] Elipsed Time: %s", time.Since(start))
}
