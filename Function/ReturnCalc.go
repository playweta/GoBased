package main

import "fmt"

func add(a, b int) (c int) {
	c = a + b
	// 命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}
func main() {
	var a, b int = 1, 2
	// 多返回值可直接作为其他函数调用实参。
	add(calc(a, b))
	c := add(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
}
