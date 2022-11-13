package main

import "fmt"

func testNil() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	testNil()
}
