package main

import "fmt"

func fooRe() (i int) {

	i = 0
	defer func() {
		fmt.Println(i)
	}()
	return 2
}

func main() {
	fooRe()
}
