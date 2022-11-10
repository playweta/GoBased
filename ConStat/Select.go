package main

import (
	"fmt"
	"time"
)

/*
1、select可以监听channel的数据流动
2、select的用法与switch语法非常类似，
由select开始的一个新的选择块，每个选择条件由case语句来描述
3、与switch语句可以选择任何使用相等比较的条件相比，
select由比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作
*/

/*
每个case都必须是一个通信
所有channel表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行；其他被忽略。
如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
否则：
如果有default子句，则执行该语句。
如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
*/
// 典型用法
// 1、超时判断
//比如在下面的场景中，使用全局resChan来接受response，如果时间超过3S,resChan中还没有数据返回，则第二条case将执行
var resChan = make(chan int)

// do request
func init() {
	select {
	case data := <-resChan:
		doData(data)

	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}
}
func doData(data int) {
	//...
	println("doData方法")
}

// 3、判断channel是否阻塞
// 在某些情况下是存在不希望channel缓存满了的需求的，可以用如下方法判断
func init() {
	ch := make(chan int, 5)
	// ...
	data := 0
	select {
	case ch <- data:
	default:
		//做相应操作，比如丢弃data。视需求而定
	}
}

// 基本使用
func init() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, "to c2\n")

	case i3, ok := <-c3:
		if ok {
			fmt.Printf("received ", i3, "from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication")
	}
}

func main() {

}
