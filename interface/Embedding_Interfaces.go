package main

import "fmt"

type ABC interface {
	abc()
}

type ZXC interface {
	zxc()
}

type Marge interface {
	ABC
	ZXC
}

type Edge struct {
	x int
}

func (e Edge) abc() {
	fmt.Println(e.x + 100)
}

func (e Edge) zxc() {
	fmt.Println(e.x - 100)
}

func main() {
	edge := Edge{10}
	var marge Marge = edge

	marge.abc()
	marge.zxc()
}
