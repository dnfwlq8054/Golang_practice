package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	startTime := time.Now()

	for i := 0; i < 100; i++ {
		arr = append(arr, i)
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("실행시간: %s\n", elapsedTime.Seconds())

	println(&arr)
	startTime = time.Now()
	arr = append(arr[:5], arr[6:]...)
	elapsedTime = time.Since(startTime)
	fmt.Printf("실행시간: %s\n", elapsedTime.Seconds())
	fmt.Println(arr)

}
