package main

func main() {
	testDeferc(0)
	/*返回
	c
	b
	a
	panic: runtime error: integer divide by zero
	*/
}

func testDeferc(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()
	defer println("c")
}
