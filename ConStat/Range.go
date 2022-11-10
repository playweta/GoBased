package main

import "fmt"

func init() {
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s {
		if i == 0 {
			s = s[:3]
			s[2] = 100
		}
		println(i, v)
	}
}

func init() {
	a := [3]int{0, 1, 2}

	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 888
			fmt.Println(a)
		}
		a[i] = v + 100
	}
	fmt.Println(a)
}

func init() {
	s := "abc"
	for i := range s {
		println(s[i])
	}
	for _, c := range s {
		println(c)
	}
	for range s {

	}
	m := map[string]int{"a": 1, "b": 2}

	for k, v := range m {
		println(k, v)
	}
}

func main() {

}
