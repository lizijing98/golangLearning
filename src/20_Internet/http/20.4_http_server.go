package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由
	// xxx/user ==> func1()
	// xxx/name ===> func2()
	// func()是回调函数，用于路由的响应，回调函数的原型是固定的
	// https:127.0.0.1/9200/user
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request ==> 客户端发来的数据
		fmt.Println("用户请求详情:", request)
		//writer ==> 通过 writer 将数据返回给客户端
		_, _ = io.WriteString(writer, "这是 /user 的请求返回数据")
	})
	// https:127.0.0.1/9200/name
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是 /name 的请求返回数据")
	})
	// https:127.0.0.1/9200/id
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是 /id 的请求返回数据")
	})

	fmt.Println("http server is running ......")
	err := http.ListenAndServe(":9300", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe err:", err)
		return
	}
}
