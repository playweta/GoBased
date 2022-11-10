package main

import (
	"fmt"
)

func first() {

	var arr1 = [3]int{1, 2}
	var arr2 = [5]int{1, 2, 3}
	var arr3 = [...]int{1, 2, 3, 4, 5, 6, 7}
	var str = [5]string{3: "hello world", 4: "termon", 0: "zero"}

	a := [3]int{1, 2}             // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 54, 5} // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200}   // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}

	fmt.Println(arr1, arr2, arr3, str)
	fmt.Println(a, b, c, d)
}

func tow() {
	var arr0 [5][3]int
	var arr1 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	{
		a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
		b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
		fmt.Println(a, b)
	}

	fmt.Println(arr0, arr1)

}

func slice(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000

	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
}

func length(a [2]int) (int, int) {
	return len(a), cap(a)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func sumArr(arr [10]int) int {
	var sum int = 0
	for v := range arr {
		sum += v
	}
	return sum
}

func eleSum(a [5]int, target int) {
	for i, v := range a {
		for j := i + 1; j < len(a); j++ {
			if a[j]+v == target {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func main() {
	//first()
	//tow()
	//a := [2]int{}
	//slice(a)
	//fmt.Println(a)
	//fmt.Println(length(a))
	//var arr1 [5]int
	//printArr(&arr1)
	//fmt.Println(arr1)
	//
	//arr2 := [...]int{2, 4, 6, 8, 10}
	//printArr(&arr2)
	//fmt.Println(arr2)
	//rand.Seed(time.Now().Unix())
	//var b [10]int
	//for i := 0; i < len(b); i++ {
	//	// 产生一个0到1000随机数
	//	b[i] = rand.Intn(1000)
	//}
	//sum := sumArr(b)
	//fmt.Printf("sum=%d\n", sum)

	b := [5]int{1, 3, 5, 8, 7}
	eleSum(b, 8)
}
