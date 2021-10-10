package main

import (
	"fmt"
	"net"
	"strings"
)

// 只能接收一个连接，只能发送一个数据
func main() {
	//1.创建监听
	ip := "127.0.0.1"
	port := 9200
	address := fmt.Sprintf("%s:%d", ip, port) // Sprintf 把 ip 和 port 拼成一个字符串
	//简写：net.Listen("tcp",":8848) 冒号前默认是 localhost
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("监听中...")

	/**
	需求：
	1.server 可以接收多个连接；==> 主 goroutine 负责监听，子 goroutine 负责数据处理
	2.每个连接可以接受处理多轮数据
	*/

	for true {
		//主 goroutine 负责监听
		//2.建立与 Client 的连接
		conn, err := listener.Accept() //只有一个 err 变量进行复用
		if err != nil {
			fmt.Println("listener.Accept() err:", err)
			return
		}
		fmt.Println("连接建立成功")

		//子 goroutine 负责数据处理
		go handleFunc(conn)
	}
}

//处理具体业务的逻辑，需要将 conn 传递进来，每一个新连接，conn 是不同的
func handleFunc(conn net.Conn) {
	for true { // 保证每个连接可以多次接受 Client 发送的数据
		//3.创建一个容器，用于接受读取到的数据
		buf := make([]byte, 1024) //使用 make 创建切片，byte <==> uint8
		fmt.Println("准备读取 Client 发来的数据")
		cnt, err := conn.Read(buf) //cnt：真正读取 client 发来的数据的长度
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("Server <=== Client,长度:", cnt, "数据", string(buf[:cnt]))

		//4.将数据转成大写 "hello" ==> "HELLO"
		upperData := strings.ToUpper(string(buf[:cnt]))

		//5.向回写入数据
		cnt, err = conn.Write([]byte(upperData))
		fmt.Println("Server ===> Client,长度:", cnt, "数据", upperData)
	}
}
