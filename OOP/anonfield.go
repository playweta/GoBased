package main

import "fmt"

// go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

// Person 人
type Person struct {
	name string
	sex  string
	age  int
}

// Student 学生
type Student struct {
	Person
	id   int
	addr string
}

type StudentZ struct {
	*Person
	id   int
	addr string
}

// 指针类型匿名字段
func init() {
	s1 := StudentZ{&Person{"5lmh", "man", 18}, 1, "bj"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.Person.name)
}

// 所有的内置类型和自定义类型都是可以作为匿名字段去使用
func init() {
	s1 := Student{Person{"5lmh", "man", 18}, 1, "bj"}
	fmt.Println(s1)
}

// 同名字段的情况
func init() {
	var s Student
	// 给自己字段赋值了
	s.name = "5lmh"
	fmt.Println(s)

	// 若给父类同名字段赋值，如下
	s.Person.name = "枯藤"
	fmt.Println(s)
}

// go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段
func main() {
	// 初始化
	s1 := Student{Person{"5lmh", "man", 20}, 1, "bj"}
	fmt.Println(s1)

	s2 := Student{Person: Person{"5lmh", "man", 20}}
	fmt.Println(s2)

	s3 := Student{Person: Person{name: "5lmh"}}
	fmt.Println(s3)
}
