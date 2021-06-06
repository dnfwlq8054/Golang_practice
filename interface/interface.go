package main

import "fmt"

type Calculator interface {
	plus() int
	minus() int
}

//Parameter 정의
type Parameter struct {
	n1, n2 int
}

//Parameter 타입에 대한 Calculator 인터페이스 구현
func (r *Parameter) plus() int {
	return r.n1 + r.n2
}

func (r *Parameter) minus() int {
	return r.n1 - r.n2
}

func showCalc(calc Calculator) {
	fmt.Println(calc.plus())
	fmt.Println(calc.minus())
}

func main() {
	r := &Parameter{30, 20}
	showCalc(r)
}
