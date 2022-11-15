package main

import "fmt"

type UserF struct {
	id   int
	name string
}

type Manager struct {
	UserF
	title string
}

func (self *UserF) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

func main() {
	m := Manager{UserF{1, "Tom"}, "Administrator"}
	//fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
	fmt.Println(m.UserF.ToString())
}
