package main

import "fmt"

type A struct {
	name string
	num  int
}

func newA() *A {
	a := A{}
	a.name = "hwan"
	a.num = 10

	return &a
}

func (a *A) print() {
	fmt.Println(a)
}

func main() {
	a := newA()
	b := A{"zzzz", 100}
	c := new(A)

	c.name = "qqqq"
	c.num = 2222

	a.print()      //output : &{hwan 10}
	fmt.Println(b) //output : {zzzz 100}
	fmt.Println(c) //output : &{qqqq 2222}
}
