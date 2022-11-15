package main

import "fmt"

type UserE struct {
	id   int
	name string
}

func (self *UserE) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func (self *UserE) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}

func (self UserE) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

func main() {
	u := UserE{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue() // 隐式传递 receiver

	mExpression := (*UserE).Test
	mExpression(&u) // 显式传递 receiver
}

func init() {
	u := UserE{1, "Tom"}
	mValue := u.Test // 立即复制 receiver，因为不是指针类型，不受后续修改影响。

	u.id, u.name = 2, "Jack"
	u.Test()

	mValue()
}

func init() {
	u := UserE{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &u, u)

	mv := UserE.TestValue
	mv(u)

	mp := (*UserE).TestPointer
	mp(&u)

	mp2 := (*UserE).TestValue // *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
	mp2(&u)
}

type DataE struct{}

func (DataE) TestValue() {}

func (*DataE) TestPointer() {}

func init() {
	var p *DataE = nil
	p.TestPointer()

	(*DataE)(nil).TestPointer() // method value
	(*DataE).TestPointer(nil)   // method expression

	// p.TestValue()            // invalid memory address or nil pointer dereference

	// (Data)(nil).TestValue()  // cannot convert nil to type Data
	// Data.TestValue(nil)      // cannot use nil as type Data in function argument
}
