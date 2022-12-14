package main

import "fmt"

func testSlice(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

func main() {
	s := []int{1, 2, 3}
	res := testSlice("sum: %d", s...) // slice... 展开slice
	println(res)
}
