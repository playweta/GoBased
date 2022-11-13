package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func testGC() {
	lock.Lock()
	lock.Unlock()
}

func testDefer() {
	lock.Lock()
	defer lock.Unlock()
}

func main() {
	func() {
		t1 := time.Now()

		for i := 0; i < 10000; i++ {
			testGC()
		}
		elapsed := time.Since(t1)
		fmt.Println("test elapsed: ", elapsed)
	}()
	func() {
		t1 := time.Now()

		for i := 0; i < 10000; i++ {
			testDefer()
		}
		elapsed := time.Since(t1)
		fmt.Println("test elapsed: ", elapsed)
	}()
}
