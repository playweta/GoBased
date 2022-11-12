package main

func addDefer(x, y int) (z int) {
	defer func() { // 3
		z += 100 // 4
	}()

	z = x + y // 1
	return    // 2 //5
}

func main() {
	println(addDefer(1, 2))
}
