package main

import "fmt"

func init() {
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

func init() {
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

func length(s string) int {
	println("call length")
	return len(s)
}
func init() {
	s := "abcd"
	for i, n := 0, length(s); i < n; i++ { // 避免多次调用 length 函数
		println(i, s[i])
	}
}

func init() {
	s := "abc"
	for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
		println(s[i])
	}

	n := len(s)
	for n > 0 { // 替代 while (n > 0) {}
		println(s[n]) // 替代 for (; n > 0;) {}
		n--
	}

	for { // 替代 while (true) {}
		println(s) // 替代 for (;;) {}
	}
}

func main() {

}
