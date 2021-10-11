package main

import (
	"fmt"
	"net"
)

//此 demo 未做多文件工程，只写在一个文件中，不做代码整理

//User 结构体
type User struct {
	Id   string
	Name string
	Msg  chan string
}

//全局的 map 结构存储所有的 User,上限设为 500
var allUsers = make(map[string]User, 500)

//全局的 message 通道用于接受任何人发送的消息
var message = make(chan string, 10)

var LF = "\n"

func main() {
	//创建服务器
	fmt.Println("服务器启动中......")
	listener, err := net.Listen("tcp", ":9301")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//启动全局唯一的 goroutine 负责监听 message channel 写给所有 User
	go broadcast()
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
		fmt.Printf("新连接:%s建立成功\n", conn.RemoteAddr().String())
		//启动处理业务的 goroutine
		go handler(conn)
	}
}

//处理具体业务
func handler(conn net.Conn) {
	//客户端与服务器建立连接时，会有 ip 和 port，是唯一的
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("%s:启动业务\n", clientAddr)
	//创建一个 User
	newUser := User{
		Id:   clientAddr,            //id 不会修改，作为在 map 中的 key
		Name: clientAddr,            //可以修改，会提供一个 rename 命令进行修改，建立连接时初始值与 id 相同
		Msg:  make(chan string, 10), //注意需要 make 空间
	}
	//启动 goroutine 负责将 msg 信息返回给 Client
	go writeBackToClient(&newUser, conn)
	//添加 newUser 至 allUsers 中
	allUsers[newUser.Id] = newUser
	loginInfoClient := fmt.Sprintf("用户:%s 你好", newUser.Name)
	newUser.Msg <- loginInfoClient
	//上线信息广播通知所有人
	loginInfoTo := fmt.Sprintf("[广播]用户:%s 已登录", newUser.Name)
	message <- loginInfoTo
	for {
		//读取客户端发送的数据
		buffer := make([]byte, 1024)
		cnt, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(newUser.Id, "conn.Read err:", err)
			return
		}
		fmt.Println("Server <==", newUser.Id, "| cnt:", cnt, "| data:", string(buffer[:cnt-1])) //读取的长度 cnt 包括了从键盘输入的换行
	}
}

//向所有用户广播消息，启动一个全局唯一的 goroutine
func broadcast() {
	fmt.Println("广播 goroutine 启动!")
	defer fmt.Println("广播 goroutine 停止")

	for {
		fmt.Println("广播 broadcast 监听中...")
		//1.从 message channel 中读取数据
		info := <-message
		fmt.Println("|广播| message channel <==", info)
		//2.将数据写入到每个用户的 msg channel 中
		for _, user := range allUsers {
			//如果 msg 是非缓冲的，会在此处会阻塞
			user.Msg <- info
		}
	}
}

//每个连接需要再启动一个 goroutine：读取 message 的数据发送给 User 的 Client
func writeBackToClient(user *User, conn net.Conn) {
	fmt.Printf("user:%s的 goroutine 正在监听自己的 msg channel\n", user.Name)
	//不断读取自身的 msg channel
	for data := range user.Msg {
		fmt.Printf("|%s| <== %s\n", user.Name, data)
		_, err := conn.Write([]byte(data + LF))
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
	}
}
