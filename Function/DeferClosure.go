package main

import "fmt"

func main() {
	var whatever [5]struct{}
	for i := range whatever {
		// i := i	// 4 3 2 1 0
		defer func() {
			fmt.Println(i)
		}() // 4 4 4 4 4
	}
}
