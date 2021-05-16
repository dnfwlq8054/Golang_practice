package main

import "fmt"

func add(a int, b int) func() {
	sum := a + b
	return func() {
		fmt.Println(sum)
	}
}

func main() {
	add(10, 20) //output : 30

	print_add := add(10, 20)
	print_add()
}
