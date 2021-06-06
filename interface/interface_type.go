package main

import (
	"fmt"
	"reflect"
)

func typeCheck(x interface{}) {
	fmt.Println(reflect.TypeOf(x))
}

func main() {
	var x interface{}
	x = 1

	fmt.Println(x)
	typeCheck(x)

	x = "Tom"
	fmt.Println(x)
	typeCheck(x)
}
