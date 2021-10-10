package main

import (
	"fmt"
	"net"
)

//此 demo 未做多文件工程，只写在一个文件中，不做代码整理

func main() {
	//创建服务器
	fmt.Println("服务器启动中......")
	listener, err := net.Listen("tcp", ":9301")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("服务器启动成功!")
	for true {
		fmt.Println("主 goroutine 监听中......")
		//监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
			return
		}
		//建立连接
		fmt.Println("连接建立成功")
		//启动处理业务的 goroutine
		go handler(conn)
	}
}

//处理具体业务
func handler(conn net.Conn) {
	for true {
		fmt.Println("启动业务")
		//读取客户端发送的数据
		buffer := make([]byte, 1024)
		cnt, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("Server <== Client | cnt:", cnt, "data:", string(buffer[:cnt-1])) //读取的长度 cnt 包括了从键盘输入的换行
	}
}
