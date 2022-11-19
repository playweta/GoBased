package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http://127.0.0.1:8000/go
	// 单独写回调函数
	http.HandleFunc("/go", myHandler)
	//http.HandleFunc("/ungo", myHandler2)
	// addr: 监听地址
	// handler: 回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)

}

func myHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method", request.Method)
	// /go
	fmt.Println("url: ", request.URL.Path)
	fmt.Println("header: ", request.Header)
	fmt.Println("body: ", request.Body)
	writer.Write([]byte("www.51mh.com"))
}
