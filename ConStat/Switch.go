package main

import "fmt"

// Type Switch
func init() {

	// 写法一
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T\r\n", i)
	case int:
		fmt.Printf("x 是 int类型")
	case float64:
		fmt.Printf("x 是 float64 类型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 类型")
	case bool, string:
		fmt.Printf("x 是 bool 或者 string 类型")
	default:
		fmt.Printf("未知类型")
	}
	// 写法二
	var j = 0
	switch j {
	case 0:
	case 1:
		fmt.Printf("1\n")
	case 2:
		fmt.Printf("2\n")
	default:
		fmt.Println("def")
	}

	// 写法三
	var k = 0
	switch k {
	case 0:
		println("fallthrough")
		fallthrough
		/*
				Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
				而如果switch没有表达式，它会匹配true。
			    Go里面switch默认相当于每个case最后带有break，
			    匹配成功后不会自动向下执行其他case，而是跳出整个switch,
			    但是可以使用fallthrough强制执行后面的case代码。
		*/
	case 1:
		fmt.Printf("1\n")
	case 2:
		fmt.Printf("2\n")
	default:
		fmt.Printf("def\n")
	}
	// 写法四
	var m = 0
	switch m {
	case 0, 1:
		fmt.Printf("1\n")
	case 2:
		fmt.Printf("2\n")
	default:
		fmt.Printf("def\n")
	}
	// 写法五
	var n = 0
	switch {
	case m > 0 && n < 10:
		fmt.Printf("m > 0 && n < 10\n")
	case n > 10 && n < 20:
		fmt.Printf("n > 10 && n < 20\n")
	default:
		fmt.Printf("def\n")
	}
}

// switch 语句
func init() {
	/*定义局部变量*/
	var grade = "B"
	var marks = 90
	switch marks {
	case 90:
		grade = "A"
	case 60:
		grade = "B"
	case 40, 50, 30:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀！\n")
	case grade == "B":
		fmt.Printf("良好！\n")
	case grade == "C":
		fmt.Printf("一般！\n")
	default:
		fmt.Printf("差！\n")
	}
	fmt.Printf("等级为:%s\n", grade)
}

func main() {

}
