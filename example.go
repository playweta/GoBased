package main

import (
	"fmt"
	"math"
	"testing"
)

var (
	a string
	b int
	c bool
	d float32
)

var f uint

const number = 1000
const (
	pi = 3.1415926
	e  = 2.7182
)

const (
	//n1 = 111  // 111, 111, 111
	n1 = iota //  0 1 2
	n2 = 100
	n6
	//_   0 1 3
	n3 = iota
	n4
	n5 // 0, 100, 3, 4 , 5
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	aj, ad = iota + 1, iota + 2 // 1 2
	aa, ab
	ac, ag
	af, az
)

var name, sex = "java", 2

func foo() (string, int) {
	return name, sex
}

func traversalString(str string) {

	for i := 0; i < len(str); i++ {
		fmt.Printf("%v(%c)\n", str[i], str[i])
	}
	fmt.Println("------------")
	for _, un := range str {
		fmt.Printf("%v(%c)\n", un, un)
	}

}

func changeString(str string) string {

	if str == "" || str == " " {
		return "空"
	}

	arr := []byte(str)
	arr[0] = 'G'
	arr[1] = 'O'
	return string(arr)
}

func sqrtDemo(aaa int, bbb int) float64 {
	var ccc float64
	ccc = math.Abs(float64(aaa*aaa + bbb*bbb))
	return ccc
}

func init() {

	fmt.Println(sqrtDemo(2, 4))

	fmt.Println(changeString("  "))

	//str := "字符串分割"
	//traversalString(str)
	x, _ := foo()
	_, v := foo()

	fmt.Printf("常量连续声明%v, %v, %v, %v\n", n1, n2, n3, n5)
	fmt.Printf("位运算连续声明%v, %v, %v, %v, %v\n", KB, MB, GB, TB, PB)
	fmt.Println(2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2)
	fmt.Println(aj, ad)
	fmt.Println(aa, ab)
	fmt.Println(ac, ag)
	fmt.Println(af, az)
	fmt.Println("str := \"c:\\pprof\\main.exe\"")
	fmt.Printf("init上: {name:%v, sex:%v}\n", x, v)
}

func main() {
	fmt.Println("go project test")
}

func init() {
	fmt.Println("init下")
}

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("add(2, 4) error, expect:%d, actual:%d", 6, r)
	}
	t.Logf("test add succ")
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
