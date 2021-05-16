package main

import (
	"fmt"
)

func add(a int, b int) (int, string) {
	return a + b, "hi"
}

func main() {
	var a int = 10
	var b int = 20

	sum, s1 := add(a, b)
	_, s2 := add(a, b)

	fmt.Println(sum, s1)
	fmt.Println(s2)
	fmt.Println(add(a, b))
}
