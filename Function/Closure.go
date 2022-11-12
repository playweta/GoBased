package main

import "fmt"

// 闭包复制的是原对象指针，这就很容易解释延迟引用现象。
func testClo() func() {
	x := 100
	fmt.Printf("x（%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}
func init() {
	f := testClo()
	f()
	// 输出
	// x（0xc000128058) = 100
	// x (0xc000128058) = 100
}

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
func init() {
	c := a()
	c()
	c()
	c()
	a()

}

func main() {

}
