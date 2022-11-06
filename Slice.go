package main

import "fmt"

func init() {
	// 声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("s1是空的")
	} else {
		fmt.Println("s1不是空")
	}

	s2 := []int{}
	// make
	var s3 = make([]int, 0)
	fmt.Println(s1, s2, s3)

	// 初始化赋值
	var s4 = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)

	arr := [5]int{1, 2, 3}
	var s6 []int
	s6 = arr[1:5]
	fmt.Println(s6)
}

func init() {

}

func main() {

}
