package main

func closure() func() {
	var i int = 0
	return func() {
		i++
		println(i)
	}
}

func main() {
	next := closure()

	next()
	next()
	next()
}
