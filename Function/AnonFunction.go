package main

import (
	"fmt"
	"math"
)

func init() {
	fn := func() { println("Hello, World!") } // 1
	fn()                                      // 2

	fns := []func(x int) int{
		func(x int) int {
			return x + 1 // 3
		},
		func(x int) int {
			return x + 2
		},
	}
	println(fns[0](100)) // 4

	// --- function as field ---

	d := struct {
		fn func() string // 5
	}{
		fn: func() string {
			return "Hello World!" // 6
		},
	}
	println(d.fn()) // 7

	fc := make(chan func() string, 2) // 8
	fc <- func() string {
		return "Hello, World" // 9
	}
	println((<-fc)()) // 10
}

func main() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))
}
