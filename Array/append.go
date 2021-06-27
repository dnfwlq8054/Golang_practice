package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr2 := []int{10, 20, 30}

	arr = append(arr, 100)               // output : [1, 2, 3, 4, 5, 100]
	arr3 := append(arr, arr2[0], arr[1]) // output : [1, 2, 3, 4, 5, 100, 10, 20]
	arr4 := append(arr, arr2...)         // output : [1, 2, 3, 4, 5, 100, 10, 20, 30]
	arr5 := append(arr[:2], arr[3:]...)  // output : [1, 2, 4, 5, 100]
	fmt.Println(arr)                     //output : [1, 2, 4, 5, 100, 100]
}
