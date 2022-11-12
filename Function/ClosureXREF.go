package main

import "fmt"

func addXREF(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}
func main() {
	temp1 := addXREF(10)
	fmt.Println(temp1(1), temp1(2))
	// 此时tmp1和tmp2不是一个实体了
	temp2 := addXREF(100)
	fmt.Println(temp2(1), temp2(2))
}
