package main

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, "closed")
}

// Close 加上他返回
//c closed
//b closed
//a closed
// 不加返回三个 c closed
func Close(t Test) {
	t.Close()
}
func main() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		//defer Close(t)返回三个 c closed
		t2 := t
		defer t2.Close()
	}
}
