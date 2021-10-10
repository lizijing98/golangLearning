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
	address := fmt.Sprintf("%s:%d", ip, port)
	//简写：net.Listen("tcp",":8848) 冒号前默认是 localhost
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("监听中...")
	conn, err := listener.Accept() //只有一个 err 变量进行复用
	if err != nil {
		fmt.Println("listener.Accept() err:", err)
		return
	}
	fmt.Println("连接建立成功")

	//2.创建一个容器，用于接受读取到的数据
	buf := make([]byte, 1024)  //使用 make 创建切片，byte <==> uint8
	cnt, err := conn.Read(buf) //cnt：真正读取 client 发来的数据的长度
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("Server <=== Client,长度:", cnt, "数据", string(buf[:cnt]))

	//3.将数据转成大写 "hello" ==> "HELLO"
	upperData := strings.ToUpper(string(buf[:cnt]))

	//4.向回写入数据
	cnt, err = conn.Write([]byte(upperData))
	fmt.Println("Server ===> Client,长度:", cnt, "数据", upperData)

	//5.关闭连接
	err = conn.Close()
	if err != nil {
		fmt.Println("conn.Close err:", err)
	}

}
