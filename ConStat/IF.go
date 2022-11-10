package main

// if	 不支持三元操作符(三目运算符) "a > b ? a : b"。
func init() {
	x := 0
	if n := "abc"; x > 0 {
		println(n[2])
	} else if x < 0 {
		println(n[1])
	} else {
		println(string(n[0]))
	}
}

func main() {

}
