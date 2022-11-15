package main

import (
	"errors"
	"fmt"
)

func getCircleAreaE(radius float32) (area float32, err error) {
	if radius < 0 {
		// 构建个异常对象
		err = errors.New("半径不能为负")
		return
	}
	area = 3.14 * radius * radius
	return
}

func main() {
	area, err := getCircleAreaE(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
}
