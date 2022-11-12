package main

import "fmt"

func test(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func main() {
	println(test(1, 2, "str"))
}
