package main

import (
	"fmt"
	"reflect"
)

func typeSwitch(n interface{}) {
	switch n.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("num")
	default:
		fmt.Println(reflect.TypeOf(n))
	}
}

func main() {
	var x interface{} = 1

	i := x
	j := x.(int)
	z, flag := x.(string)

	typeSwitch(i)
	typeSwitch(j)
	typeSwitch(z)
	typeSwitch(flag)

}
