package main

import "fmt"

func swap(x, y *int) {
	var temp int

	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var a, b int = 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
