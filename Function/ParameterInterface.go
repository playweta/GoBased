package main

import "fmt"

func testPar(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

func main() {
	println(testPar("sunm: %d", 1, 2, 3))
}
