package main

type stru struct {
	name string
	num  int
}

func (b *stru) Sound() {
	println("qauck")
	println(b.name, b.num)
}

type Sounder interface {
	Sound()
}

func main() {
	var d stru = stru{"1", 2}
	d.Sound()
	//fmt.Println(d)
}
