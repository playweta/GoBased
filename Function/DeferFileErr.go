package main

import (
	"fmt"
	"io"
	"os"
)

func doFile() error {
	file, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if file != nil {
		defer func(file io.Closer) {
			if err := file.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}(file)
	}
	file, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if file != nil {
		defer func(file io.Closer) {
			if err := file.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}(file)
	}
	return nil
}

func main() {
	println(doFile())
}
