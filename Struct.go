package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// 实现map有序输出(面试经常问
func init() {
	mapUrl := make(map[int]string, 5)
	mapUrl[1] = "www.topgoer.com"
	mapUrl[2] = "rpc.topgoer.com"
	mapUrl[4] = "JDBC.DATA"
	mapUrl[5] = "xiaohong"
	mapUrl[3] = "xiaohuang"
	var sli []int
	for k, _ := range mapUrl {
		sli = append(sli, k)
	}
	sort.Ints(sli)
	for i := 0; i < len(mapUrl); i++ {
		fmt.Println(mapUrl[sli[i]])
	}

}

// 删除map类型的结构体
func init() {
	ce := make(map[int]student)
	ce[1] = student{1, "小明哥", 23}
	ce[2] = student{2, "万豪", 56}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}

// 小练习

func demo(ce []student) {
	// 切片是引用传递，是可以改变值的
	ce[1].age = 999
}
func init() {
	var ce []student // 定义一个切片类型的机构体
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 54},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
}

// 结构体标签（Tag）
func init() {
	student := Student{
		ID:     1,
		Gender: "女",
		Name:   "标签员",
	}
	data, err := json.Marshal(student)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data)
}

// 结构体与JSON序列化

// Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	Name   string //私有不能被json包访问
}

// Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func init() {
	class := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu%02d", i),
		}
		class.Students = append(class.Students, stu)
	}
	// JSON序列化：结构体--》JSON格式化的字符串
	data, err := json.Marshal(class)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	// json反序列化
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed")
		return
	}
	fmt.Printf("%#v\n", c1)
}

// 结构体的“继承”

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动!\n", a.name)
}

// Dog 狗
type Dog struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会旺旺旺~\n", d.name)
}
func init() {
	dog := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	dog.wang()
	dog.move()
}

// 嵌套结构体的字段名冲突

// Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

func init() {
	var user User
	user.Name = "汪峰"
	user.Gender = "男"
	user.Address.CreateTime = "2000"
	user.Email.CreateTime = "2000"
}

// 嵌套匿名结构体
func init() {
	var user User
	user.Name = "元亓"
	user.Gender = "无"
	user.Address.Province = "陕西" //通过匿名结构体.字段名访问
	user.Address.City = "西安"     //直接访问匿名结构体的字段名
	fmt.Printf("user=%#v\n", user)
}

// 嵌套结构体

// Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

// User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
	Email
}

func init() {
	user := User{
		Name:   "olo",
		Gender: "男",
		Address: Address{
			Province: "shanghai",
			City:     "xizhang",
		},
	}
	fmt.Printf("user=%#v\n", user)
}

// 结构体的匿名字段

// Person 结构体Person类型
type Person struct {
	string
	int
}

func init() {
	per := Person{
		"娃娃鱼",
		45,
	}
	fmt.Printf("%#v\n", per)
	fmt.Println(per.string, per.int)
}

// 任意类型添加方法

// MyInt 将int定义为自定义MyInt类型
type MyInt int

// SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int")
}
func init() {
	var my MyInt
	my.SayHello()
	my = 100
	fmt.Printf("%#v\t %T\n", my, my)
}

// 值类型的接收者

// SetAgeVal 设置p的年龄
// 使用值接收者
func (per Per) SetAgeVal(newAge int8) {
	per.age = newAge
}
func init() {
	per := newPer("杀猪菜", 25)
	per.Dream()
	fmt.Printf("年龄=》%v\n", per.age) // 25
	per.SetAgeVal(30)
	fmt.Printf("最新年龄=》%v\n", per.age) //25
}

// 指针类型的接收者
// 1.需要修改接收者中的值
// 2.接收者是拷贝代价比较大的大对象
// 3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

// SetAge 设置p的年龄
// 使用指针接收者
func (per *Per) SetAge(newAge int8) {
	per.age = newAge
}
func init() {
	per := newPer("菜刀", 35)
	fmt.Println(per.age)
	per.SetAge(40)
	fmt.Println(per.age)
}

// 方法和接收者

// Per 结构体
type Per struct {
	name string
	age  int8
}

// newPer 构造函数
func newPer(name string, age int8) *Per {
	return &Per{
		name: name,
		age:  age,
	}
}

// Dream Per的方法
func (per Per) Dream() {
	fmt.Printf("%s的梦想是学习好Goyuyan！\n", per.name)
}

func init() {
	per := newPer("小娜", 35)
	per.Dream()
}

// 构造函数
func newPerson(name, city string, age int8) *person {
	return &person{name: name, city: city, age: age}
}
func init() {
	user := newPerson("dod", "test", 90)
	fmt.Printf("%#v\n", user)
}

//  面试题 --结构体
type student struct {
	id   int
	name string
	age  int
}

func init() {
	m := make(map[string]*student)
	stus := []student{
		{name: "张三", age: 34},
		{name: "王五", age: 30},
		{name: "李四", age: 39},
	}
	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

// 结构体内存布局
func init() {
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}

	n := test{
		1, 2, 3, 4,
	}

	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}

// 使用值的列表初始化
//  1.必须初始化结构体的所有字段。
//  2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
//  3.该方式不能和键值初始化方式混用。
func init() {
	user := &person{
		"java",
		"sum",
		20,
	}

	fmt.Printf("user=%#v\n", user)
}

// 对结构体指针进行键值对初始化
func init() {
	user := &person{
		name: "user",
		city: "nanjing",
		age:  9,
	}
	fmt.Printf("user=%#v\n", user)
}

// 使用键值对初始化
func init() {
	user := person{
		name: "liu",
		city: "jianxi",
		age:  1,
	}
	fmt.Printf("user=%#v\n", user)
}

// 结构体初始化
func init() {
	var user person
	fmt.Printf("user=%#v\n", user)
}

// 取结构体的地址实例化
func init() {
	user := &person{}
	fmt.Printf("%T\n", user)
	fmt.Printf("user=%#v\n", user)
	user.name = "地址"
	user.city = "main"
	user.age = 9
	fmt.Printf("user=%v\n", user)
}

// 创建指针类型结构体
func init() {
	var user = new(person)
	user.name = "指针"
	user.city = "指向值"
	user.age = 0
	fmt.Printf("%#v\n", user) // &main.person{name:"指针", city:"指向值", age:0}
}

// 匿名结构体
func init() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "方糖"
	user.Age = 2
	fmt.Printf("%#v\n", user)
}

// 基本实例化
type person struct {
	name string
	city string
	age  int8
}

func init() {
	var root person
	root.name = "管理员"
	root.city = "江西"
	root.age = 19
	fmt.Printf("root=%v\n", root)  //p1={pprof.cn 北京 18}
	fmt.Printf("root=%#v\n", root) //p1=main.person{name:"pprof.cn", city:"北京", age:18}
}

// 类型定义和类型别名的区别
func init() {
	//类型定义
	type NewInt int

	// 定义别名
	type MyInt = int

	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) // type of a:main.NewInt
	fmt.Printf("type of a:%T\n", b) // type of b:int
}

func main() {

}
