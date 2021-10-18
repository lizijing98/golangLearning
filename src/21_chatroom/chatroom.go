package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

//此 demo 未做多文件工程，只写在一个文件中，不做代码整理

//User 结构体
type User struct {
	Id   string
	Name string
	Msg  chan string
}

var lock sync.RWMutex

//全局的 map 结构存储所有的 User,上限设为 500
var allUsers = make(map[string]User, 500)

//全局的 message 通道用于接受任何人发送的消息
var message = make(chan string, 10)

var LF = "\n"

func main() {
	//创建服务器
	fmt.Println("|Server log|服务器启动中......")
	listener, err := net.Listen("tcp", ":9301")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//启动全局唯一的 goroutine 负责监听 message channel 写给所有 User
	go broadcast()
	fmt.Println("|Server log|服务器启动成功!")
	for true {
		fmt.Println("|Server log|主 goroutine 监听中......")
		//监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen.Accept err:", err)
			return
		}
		//建立连接
		fmt.Printf("|Server log|新连接:%s建立成功\n", conn.RemoteAddr().String())
		//启动处理业务的 goroutine
		go handler(conn)
	}
}

//处理具体业务
func handler(conn net.Conn) {
	//客户端与服务器建立连接时，会有 ip 和 port，是唯一的
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("|Server log|%s:启动业务\n", clientAddr)
	//创建一个 User
	newUser := User{
		Id:   clientAddr,            //id 不会修改，作为在 map 中的 key
		Name: clientAddr,            //可以修改，会提供一个 rename 命令进行修改，建立连接时初始值与 id 相同
		Msg:  make(chan string, 10), //注意需要 make 空间
	}
	//启动 goroutine 负责将 msg 信息返回给 Client
	go writeBackToClient(&newUser, conn)
	//添加 newUser 至 allUsers 中
	lock.Lock()
	allUsers[newUser.Id] = newUser
	lock.Unlock()
	//定义 isQuit channel 用于处理退出信号
	var isQuit = make(chan bool)
	//定义 resetTimer channel 用于告知 clear() 保活用户
	var resetTimer = make(chan bool)
	//启动 goroutine 负责监听退出信号
	go clear(&newUser, conn, isQuit, resetTimer)
	loginInfoClient := fmt.Sprintf("用户:%s 你好", newUser.Name)
	newUser.Msg <- loginInfoClient
	//上线信息广播通知所有人
	loginInfoTo := fmt.Sprintf("[广播]用户:%s 已登录", newUser.Name)
	message <- loginInfoTo
	//具体业务
	for {
		//读取客户端发送的数据
		buffer := make([]byte, 1024)
		cnt, err := conn.Read(buffer)
		//如果读到的数据的 cnt==0 则为 ^C 主动退出
		//TODO:此处超时退出会触发，需要优化，修改主动退出逻辑，利用 os.Signal 包中信号捕获实现退出
		if cnt == 0 {
			fmt.Println(newUser.Id, "conn.Read err:", err, "cnt:", cnt)
			fmt.Printf("|Server log|[%s]客户端主动退出(^C),开始进行进行清理工作\n", newUser.Id)
			//在这里不进行真正的退出动作，而是发送一个退出信号，统一做退出处理，使用一个新的 channel 做信号传递
			isQuit <- true
		}
		if err != nil {
			fmt.Println(newUser.Id, "conn.Read err:", err, "cnt:", cnt)
			return
		}
		receivedStr := string(buffer[:cnt-1])
		fmt.Println("|Server log|Server <==", newUser.Id, "| cnt:", cnt, "| data:", receivedStr) //读取的长度 cnt 包括了从键盘输入的换行
		//======业务逻辑开启======
		switch {
		//1.查询当前所有用户命令 -who
		//  a.先判断接收的数据是不是 -who ==> 长度==4&&字符串=="-who"
		//  b.遍历 allUsers
		case len(receivedStr) == 4 && receivedStr == "-who":
			toClient := "所有在线用户:"
			lock.Lock()
			for _, user := range allUsers {
				toClient = fmt.Sprintf("%s\nid:%s username:%s", toClient, user.Id, user.Name)
			}
			lock.Unlock()
			newUser.Msg <- toClient
			resetTimer <- true
		//2.用户重命名 -rename|[new_name]
		// a.先判断接受的数据是不是重命名命令 ==> -rename 开头
		// b.取接受的数据中 -rename 后的 new_name 作为新的用户名
		// c.更新 newUser 的 Name 字段并更新 allUsers 中的数据
		case strings.HasPrefix(receivedStr, "-rename|"):
			newName := strings.Split(receivedStr, "|")[1]
			var result string
			if newName != "" {
				newUser.Name = newName
				lock.Lock()
				allUsers[newUser.Id] = newUser //更新 allUsers 中的数据
				lock.Unlock()
				result = fmt.Sprintf("[%s]:当前用户名为:%s", newUser.Name, newUser.Name)
				newUser.Msg <- result
			} else {
				result = fmt.Sprintf("[%s]:新用户名不能为空", newUser.Name)
			}
			newUser.Msg <- result
			resetTimer <- true
		//3.用户主动退出 -exit
		case len(receivedStr) == 5 && receivedStr == "-exit":
			fmt.Printf("|%s|客户端主动退出(-exit),开始进行进行清理工作\n", newUser.Id)
			isQuit <- true
		default:
			userSend := fmt.Sprintf("[%s]:%s", newUser.Name, receivedStr)
			message <- userSend
			resetTimer <- true
		}
		//======业务逻辑结束======
	}
}

//向所有用户广播消息，启动一个全局唯一的 goroutine
func broadcast() {
	fmt.Println("|Server log|广播 goroutine 启动!")
	defer fmt.Println("|Server log|广播 goroutine 停止")

	for {
		fmt.Println("|Server log|广播 broadcast 监听中...")
		//1.从 message channel 中读取数据
		info := <-message
		fmt.Printf("|广播| message channel <==\"%s\"\n", info)
		//2.将数据写入到每个用户的 msg channel 中
		lock.Lock()
		for _, user := range allUsers {
			//如果 msg 是非缓冲的，会在此处会阻塞
			user.Msg <- info
		}
		lock.Unlock()
	}
}

//每个连接需要再启动一个 goroutine：读取 message 的数据发送给 User 的 Client
func writeBackToClient(user *User, conn net.Conn) {
	fmt.Printf("|%s|user:%s的 goroutine 正在监听自己的 msg channel\n", user.Id, user.Name)
	//不断读取自身的 msg channel
	for data := range user.Msg {
		fmt.Printf("|%s| <== \"%s\"\n", user.Id, data)
		_, err := conn.Write([]byte(data + LF))
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
	}
}

//启动一个 goroutine 只负责监听退出信号，触发后进行清理工作：delete map、close conn
func clear(user *User, conn net.Conn, isQuit, resetTimer chan bool) {
	fmt.Printf("|%s|'退出监听' goroutine 已启动\n", user.Id)
	defer fmt.Printf("|%s|'退出监听' goroutine 已退出\n", user.Id)
	for {
		select {
		case <-isQuit:
			fmt.Printf("|%s|清理工作进行中...\n", user.Id)
			lock.Lock()
			delete(allUsers, user.Id)
			lock.Unlock()
			_ = conn.Close()
			logoutInfo := fmt.Sprintf("[广播]:用户%s已退出", user.Name)
			message <- logoutInfo
			return
		case <-time.After(60 * time.Second):
			fmt.Printf("|Server log|[%s]客户端超时退出,开始进行清理工作", user.Id)
			fmt.Printf("|%s|清理工作进行中...\n", user.Id)
			lock.Lock()
			delete(allUsers, user.Id)
			lock.Unlock()
			_ = conn.Close()
			logoutInfo := fmt.Sprintf("[广播]:用户%s已超时离线", user.Name)
			message <- logoutInfo
			return
		case <-resetTimer:
			fmt.Printf("|%s|计时器已重置\n", user.Id)
		}
	}
}
