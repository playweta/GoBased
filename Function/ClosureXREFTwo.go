package main

import "fmt"

func test01(base int) (func(int) int, func(int) int) {
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

func main() {
	f1, f2 := test01(10)
	fmt.Println(f1(1), f2(2)) // 9
	fmt.Println(f1(3), f2(4)) // 8
}
