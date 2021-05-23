package main

func nextValue() func() int {
	i := 0
	println("함수 ", &i)
	return func() int {
		i++
		println("클로저 ", &i)
		return i
	}
}

func main() {
	i := 10
	next := nextValue()
	println(i)
	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := nextValue()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2
}
