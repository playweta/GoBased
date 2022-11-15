package main

import "fmt"

type S struct {
	T
}

type T struct {
	int
}

func (t T) testT() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法")
}
func (t *T) testP() {
	fmt.Println("类型 *T 方法集包含全部 receiver *T 方法。")
}

func main() {
	s1 := S{T{1}}
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	s1.testP()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()
	s2.testP()
}
