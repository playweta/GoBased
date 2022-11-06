package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//下面的代码演示了map中值为切片类型的操作：
func init() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	val, ok := sliceMap[key]
	if !ok {
		val = make([]string, 0, 2)
	}
	val = append(val, "北京", "上海")
	sliceMap[key] = val
	fmt.Println(sliceMap)
	fmt.Println("------------------")
}

// 下面的代码演示了切片中的元素为map类型时的操作：
func init() {
	var mapSlice = make([]map[string]string, 3)
	for key, val := range mapSlice {
		fmt.Printf("key:%d val:%v\n", key, val)
	}
	fmt.Println("after init")
	// 初始化切片
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "root"
	mapSlice[0]["username"] = "root"

	for key, val := range mapSlice {
		fmt.Printf("key:%d val:%v\n", key, val)
	}
}

// 按照指定顺序遍历map
func init() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}

// map的遍历
// 使用delete()函数删除键值对
func init() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["王五"] = 80
	scoreMap["李四"] = 70
	delete(scoreMap, "张三")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

// 判断某个键是否存在
func init() {
	scoreMap := make(map[string]int)
	scoreMap["小李"] = 20
	scoreMap["小弟"] = 30

	v, ok := scoreMap["小李"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

}

// map基本使用
func init() {
	userInfo := map[string]string{
		"username": "pprof.cn",
		"password": "12345",
	}
	fmt.Println(userInfo)
}

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["栈"] = 20
	scoreMap["明"] = 40
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["栈"])
	fmt.Printf("type of a:%T\n", scoreMap)
}
