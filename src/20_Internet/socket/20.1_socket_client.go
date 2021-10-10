package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//1.与服务器创建链接
	conn, err := net.Dial("tcp", ":9200")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	fmt.Println("Client 与 Server 的链接建立成功")

	//2.向服务器发送数据
	sendData := []byte("hello")

	for true { // 保证多次向 Server 发送数据
		cnt, err := conn.Write(sendData)
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}

		fmt.Println("Client ==> Server cnt:", cnt, "data:", string(sendData))

		//3.接收服务器返回数据
		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("Client <== Server,cnt:", cnt, "data:", string(buf[:cnt]))
		time.Sleep(1 * time.Second)
	}

	//4.关闭链接
	err = conn.Close()
	if err != nil {
		fmt.Println("conn.Close err:", err)
	}

}
