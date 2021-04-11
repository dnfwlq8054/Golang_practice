package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var i int = 1
	for i < 10 {
		i *= 2
		fmt.Println(i)
	}

	names := []string{"abc", "hwan", "banana"}

	for idx, name := range names {
		fmt.Println(idx, name)
	}
}
