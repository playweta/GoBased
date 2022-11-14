package main

import "fmt"

func testRP(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("%v / %v = %d\n", x, y, z)
}
func init() {
	testRP(2, 1)
}

func except() {
	fmt.Println(recover())
}

func testEcx() {
	defer except()
	panic("test panic")
}

func init() {
	testEcx()
}

func testRec() {
	defer func() {
		fmt.Println(recover()) // 有效
	}()

	defer recover()              // 无效
	defer fmt.Println(recover()) // 无效
	defer func() {
		println("defer inner")
		recover() // 无效
	}()

	panic("test panic")
}

func init() {
	testRec()
}

func testPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}

func init() {
	testPanic()
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var ch = make(chan int, 10)
	close(ch)
	ch <- 1

}
