package main

import "fmt"

func sort(arr []int, len int, cmp func(int, int) bool) {
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if cmp(arr[i], arr[j]) {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func main() {

	var arr = []int{5, 10, 2, 69, 122}

	sort(arr, 5, func(a int, b int) bool {
		return a > b
	})

	fmt.Println(arr)
}
