package main

import "fmt"

// 指针小练习
func init() {
	var a int
	fmt.Println(&a)

	var b *int
	b = &a
	*b = 100
	fmt.Println(a)
}

// new 和 make的区别
// 1、二者都是用来做内存分配的
// 2、make只用于slice、map以及channel的初始化，返回的还是这三个引用类型
// 3、而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针

// make
func init() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}

// new
func init() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*a)
	fmt.Println(*b)

	var c *int
	c = new(int)
	*c = 10
	fmt.Println(*c)
}

// new 和 make
/*
执行上面的代码会引发panic，为什么呢？
在Go语言中对于引用类型的变量，我们在
使用的时候不仅要声明它，还要为它分配内
存空间，否则我们的值就没办法存储。而对
于值类型的声明不需要分配内存空间，是因为
它们在声明的时候已经默认分配好了内存空间。
要分配内存，就引出来今天的new和make。
Go语言中new和make是内建的两个函数，主要用来分配内存
*/
/*func init() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}*/

// 空指针
func init() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

var u = 100

func getVal(x int) {
	x = u
}

func getAddress(x *int) {
	*x = u
}

// 指针传值示例：
func init() {
	a := 10
	getVal(a)
	fmt.Println(a)
	getAddress(&a)
	fmt.Println(a)
}

// 指针取值
func init() {
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中

	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

// 指针地址和指针类型
func init() {
	a := 10
	b := &a
	fmt.Printf("a:%d\t ptr:%p\n", a, &a)
	fmt.Printf("b:%p\t type:%T\n", b, b)
	fmt.Println(*b)
	fmt.Println(&b)
}

func main() {

}
