package main

func addED(x, y int) (z int) {
	defer func() {
		println(z) // 205
	}()
	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func main() {
	println(addED(2, 3))
}
