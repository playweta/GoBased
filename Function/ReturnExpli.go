package main

func addExpli(x, y int) (z int) {
	{ // 不能在一个级别，引发 "z redeclared in this block" 错误
		var z = x + y
		//return // z is shadowed during return
		return z
	}
}

func main() {
	println(addExpli(1, 2))
}
