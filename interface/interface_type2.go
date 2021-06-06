package main

import "fmt"

func main() {
	var x interface{} = 1

	i := x
	j := x.(int)
	z, flag := x.(string)
	y := nil
	fmt.Println(i)       // output : 1
	println(i)           // output : 주소값 출력
	println(j)           // output : 1
	fmt.Println(z, flag) // false
	println(y)
}
