package main

import "fmt"

type Data struct {
	x int
}

func (self Data) ValueTest() {
	fmt.Printf("Value：%p\n", &self)
}

func (self *Data) PointerTest() {
	fmt.Printf("Ponter: %p\n", self)
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data：%p\n", p)

	d.ValueTest()
	d.PointerTest()

	p.ValueTest()
	p.PointerTest()

}
