package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

// Sayer 接口
type Sayer interface {
	say()
}

type dog struct{}

type cat struct{}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// 值接收者和指针接收者实现接口的区别
func init() {

}

// 接口类型变量
func init() {
	var x Sayer // 声明一个Sayer类型的变量x
	a := cat{}  // 实例化一个cat
	b := dog{}  // 实例化一个dog
	x = a       // 可以把cat实例直接赋值给x
	x.say()     // 喵喵喵
	x = b       // 可以把dog实例直接赋值给x
	x.say()     // 汪汪汪
}

// 为什么要使用接口
func init() {
	c := Cat{}
	fmt.Println("猫:", c.Say())
	d := Dog{}
	fmt.Println("狗:", d.Say())
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

type Mover interface {
	move()
}

func (d dog) move() {
	fmt.Println("狗会动")
}

// 值接收者实现接口
func init() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}

// 指针接收者实现接口
func (d *dog) moveZ() {
	fmt.Println("狗会动")
}
func init() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x不可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
}

// People 下面的代码是一个比较好的面试题
type People interface {
	Speak(string) string
}

type StudentX struct{}

func (stu *StudentX) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func init() {
	var peo People = &StudentX{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

// 类型与接口的关系
type dogZ struct {
	name string
}

type car struct {
	brand string
}

// 实现Sayer接口
func (d dogZ) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dogZ) move() {
	fmt.Printf("%s会动\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

// SayerZ 接口
type SayerZ interface {
	say()
}

// MoverZ 接口
type MoverZ interface {
	move()
}

// 实现Sayer接口
func (d dogZ) sayZ() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dogZ) moveZ() {
	fmt.Printf("%s会动\n", d.name)
}

func init() {
	var x Sayer
	var y Mover

	var a = dogZ{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}

func init() {
	var x Mover
	var a = dogZ{name: "旺财"}
	var b = car{brand: "保时捷"}
	x = a
	x.move()
	x = b
	x.move()
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

//空接口的定义
func init() {
	// 定义一个空接口x
	var x interface{}
	s := "pprof.cn"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// 空接口作为map的值
func init() {
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "李白"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

// 接口值
func init() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = nil
}

func init() {
	var x interface{}
	x = "pprof.cn"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
