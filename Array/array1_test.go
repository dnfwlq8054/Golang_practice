package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkF1(b *testing.B) {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100000000; i++ {
		arr = append(arr, rand.Int())
	}

	for i := 0; i < b.N; i++ {
		F1(arr)
	}
}
