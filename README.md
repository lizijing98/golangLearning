# Golang 学习笔记

<!-- @import "[TOC]" {cmd="toc" depthFrom=2 depthTo=4 orderedList=false} -->

<!-- code_chunk_output -->

- [一、前注](#一-前注)
- [二、概述](#二-概述)
  - [1.工程概述](#1工程概述)
  - [2.GO 环境变量](#2go-环境变量)
  - [3.GO 基本语句](#3go-基本语句)
- [三、GO 基本语法](#三-go-基本语法)
  - [1.文件名、关键字和标识符](#1文件名-关键字和标识符)
  - [2.GO 不支持的语法](#2go-不支持的语法)
  - [3.slice 的容量与长度](#3slice-的容量与长度)
  - [4.字典 map](#4字典-map)
  - [5.内存逃逸](#5内存逃逸)
  - [6.iota 常量组累加器](#6iota-常量组累加器)
  - [7.init 函数](#7init-函数)
  - [8.defer 延迟](#8defer-延迟)
- [四、面向对象操作](#四-面向对象操作)
  - [1.类封装](#1类封装)
  - [2.绑定方法](#2绑定方法)
  - [3.继承](#3继承)
  - [4.多态](#4多态)
- [五、并发操作](#五-并发操作)
  - [1.前提](#1前提)
  - [2.启动 go 程](#2启动-go-程)
  - [3.提前退出 go 程](#3提前退出-go-程)
  - [4.go 程之间通信 channel](#4go-程之间通信-channel)
- [六、网络编程](#六-网络编程)
  - [1.网络分层](#1网络分层)
  - [2.socket](#2socket)
  - [3.http](#3http)
    - [①.概述](#1概述)
    - [②.http 请求报文格式](#2http-请求报文格式)
    - [③.http 响应报文格式](#3http-响应报文格式)
  - [4.JSON](#4json)
  - [5.结构体的标签](#5结构体的标签)

<!-- /code_chunk_output -->

## 一、前注

go 官方网站：[The Go Programming Language (google.cn)](https://golang.google.cn/)

go 官方文档：[Documentation - The Go Programming Language (google.cn)](https://golang.google.cn/doc/)

go 语言 SDK：[Downloads - The Go Programming Language (google.cn)](https://golang.google.cn/dl/)，或者 `brew install go`

IDE:Goland 或者 VSCode

## 二、概述

==go 语言对程序格式要求非常高==

### 1.工程概述

go 语言项目需要有特定的目录解构进行管理，一个标准的 go 工程起码有三个目录：

- src
  - 放源代码文件
- bin
  - 放编译后的程序：`go install`
- pkg
  - 放依赖包

### 2.GO 环境变量

go SDK 的安装位置：`GOROOT`

工程位置环境变量：`GOPATH`

目标机器的操作系统，在 `go build` 时会用到：`GOOS`

目标机器的处理器架构：`GOARCH` ，可以是 386、amd64、arm

### 3.GO 基本语句

```
go build //编译自身包和依赖包
go install //编译并安装自身包和依赖包
gofmt //格式化代码
go doc //生成代码文档
```

## 三、GO 基本语法

### 1.文件名、关键字和标识符

**文件名**

Go 的源文件以 `.go` 为后缀名存储，文件名均有小写字母组成，允许使用下划线，不包含空格或其他特殊字符，且必须以小写字母开头，eg：`test_1.go`

**关键字**

|  **break**   |   **default**   |  **func**  | **interface** | **select** |
| :----------: | :-------------: | :--------: | :-----------: | :--------: |
|   **case**   |    **defer**    |   **go**   |    **map**    | **struct** |
|   **chan**   |    **else**     |  **goto**  |  **package**  | **switch** |
|  **const**   | **fallthrough** |   **if**   |   **range**   |  **type**  |
| **continue** |     **for**     | **import** |  **return**   |  **var**   |

**标识符**

| **append** |  **bool**   |  **byte**   |   **cap**   | **close**  | **complex** | **complex64** | **complex128** | **uint16**  |
| :--------: | :---------: | :---------: | :---------: | :--------: | :---------: | :-----------: | :------------: | :---------: |
|  **copy**  |  **false**  | **float32** | **float64** |  **imag**  |   **int**   |   **int8**    |   **int16**    | **uint32**  |
| **int32**  |  **int64**  |  **iota**   |   **len**   |  **make**  |   **new**   |    **nil**    |   **panic**    | **uint64**  |
| **print**  | **println** |  **real**   | **recover** | **string** |  **true**   |   **uint**    |   **uint8**    | **uintptr** |

### 2.GO 不支持的语法

1. 不支持前置自增自减：++i、--i
2. 不支持地址加减
3. 不支持三目运算(?:)
4. 只有 true/false 才能代表逻辑真/假，0 和 nil 不行

### 3.slice 的容量与长度

![image-20211004114838295](https://i.loli.net/2021/10/18/itNekj7LwSgxYmr.png)

go 为了保障效率直接分配多一倍的容量，但长度还是实际的使用长度，在容量为达到当前上限前不会再分配，直至达到上限

[深入理解 Go Slice - SegmentFault 思否](https://segmentfault.com/a/1190000017341615)

### 4.字典 map

哈希表(key<==>value)，key 是经过哈希运算的

### 5.内存逃逸

![image-20211004163228714](https://i.loli.net/2021/10/18/wtgo7CxFinRDOjM.png)

```shell
go build -o 10_memory_escape --gcflags "-m -m -l" 10_memory_escape.go > output 2>&output
```

### 6.iota 常量组累加器

1. iota 是常量组计数器
2. iota 从 0 开始，每换行递增加1
3. 常量组有个特点，如果不赋值，默认与上一行表达式相同
4. 如果同一行出现两个 iota 则这两个 iota 值相同
5. 每个常量组的 iota 是独立的，iota 遇到新的 const 会重置为 0

### 7.init 函数

1. C 语言中没有 init 函数，一般自己写 init 函数再构造函数中调用
2. go 语言中再带 init 函数，每一个包中可以包含一个或多个 init 函数
3. 这个 init 函数再包被引用的时候（import）进行自动调用，不允许显式调用
4. 有时引用一个包时只想引用包中的 init 函数（eg：MySQL 的 init 对驱动进行初始化），为了避免编译器报错可以使用 '_' 避免报错：`import _ "xxx/xxx"`

[【Go语言学习】包、Init函数与执行顺序_Eric_zhang929的博客-CSDN博客](https://blog.csdn.net/Eric_zhang929/article/details/102550955)

### 8.defer 延迟

1. 关键字：defer 可以用于修饰语句、函数，确保这条语句可以在当前栈退出的时候执行

    ```c
    lock.Lock()
    a = "hello"
    lock.Unlock()//<==经常容易忘掉解锁
    ```

    go 语言中

    ```go
    {
        lock.Lock()
        defer lock.Unlock()//<==在当前栈退出的时候（例如：函数结束时）执行
        a = "hello"
    }
    
    /*例如在文件读写时*/
    {
        file1,_ := file.Open()
        defer file1.Close()
    }
    ```

2. 一般用于做资源清理工作

3. 解锁、关闭文件

4. 在同一个函数中可以多次调用 defer，类似于栈的机制：先入后出

## 四、面向对象操作

### 1.类封装

go 语言支持类的操作，但是没有 class 关键字，使用 struct 来模拟类

### 2.绑定方法

go 中类的方法不在结构体中写，在结构体外写再绑定

```go
func (变量名 类名/类指针) 方法名(参数)返回值{}
```

```go
package classSrc

import "fmt"

*type Person struct {
	name   string
	age    int
	gender string
	score  float64
}

// go 中对象的方法不在结构体中写，在结构体外写再绑定

func (p *Person) EatPointer() {
	// 类的方法可以使用自己的成员/指针
	// 使用指针操作对象本身，不使用指针操作对象的拷贝
	p.Name = "LiLi"
	fmt.Printf("%s is eating\n", p.Name)
}
func (p Person) Eat() {
	// 类的方法可以使用自己的成员/指针
	// 使用指针操作对象本身，不使用指针操作对象的拷贝
	p.Name = "LiLi"
	fmt.Printf("%s is eating\n", p.Name)
}
```

### 3.继承

==golang 中并不存在真正的继承，只是 struct 之间的组合==

### 4.多态

C++：使用纯虚函数替代接口，实现多态

go：使用接口（interface）实现多态

interface 不仅仅是用于处理多态，可以接受任意数据类型，有点类似 void

C 中的多态需要父子继承关系，go 中的多态只需要实现接口即可

## 五、并发操作

### 1.前提

并发：比如电脑同时听歌看小说玩游戏全开。CPU 根据时间片进行划分，交替执行这三个程序，人可以感觉是同时运行的

并行：多个 CPU（多核）同时执行

C 语言中实现并发使用的是多线程（C++中的最小资源单元），进程

go 语言中使用的不是线程，而是 go 程==>goroutine，go 程是 go 语言原生支持的

每一个 go 程占用的系统资源远远小于线程，一个 go 程大约需要 4k~5k 的内存资源

一个程序可以启动大量的 go 程

- 线程==>几十个
- go 程==>成百上千个，对于实现高并发，性能非常好

### 2.启动 go 程

只需要在目标函数前加上 go 关键字即可

### 3.提前退出 go 程

```go
//GOEXIT ==> 提前退出当前 go 程
//return ==> 返回当前函数
//exit ==> 退出当前进程
```

### 4.go 程之间通信 channel

**有缓存 channel**

1. 当缓冲区写满时，写阻塞，当被读取后，再恢复写入
2. 当缓冲区读取完毕，读阻塞
3. 如果管道没有使用 make 分配空间，那么管道默认是 nil 的，读取写入都会阻塞
4. 从一个 nil 的 channel 读入或写入都会造成阻塞（==注意，不会崩溃==）
5. 对于一个 channel，读与写次数必须对等，否则：
   1. 在多个 go 程中会造成资源泄露
   2. 在主 go 程中，会造成死锁

**当 channel 读写次数不一致**

1. 当发生在主 go 程，发生死锁，程序会崩溃
2. 当发生在子 go 程，会造成内存泄露
3. 避免此情况，在写入端写入完成后进行 `close()` 关闭；在读出端使用 `for:range` 进行遍历

**channel 的关闭**

1. 从一个已经 close 的 channel 读取数据时会返回零值（不会崩溃）
2. 向一个 close 的 channel 写入数据会崩溃
3. 重复关闭一个已经 close 的 channel，程序会崩溃
4. 关闭管道的动作一定要在写入方，不应该放在读出方，否则写入端的继续写会造成崩溃

**判断一个管道是否已经关闭**

```go
numsChan := make(chan int, 10)
value, isOpen := <-numsChan
```

**单向通道**

为了明确语义，一般用于函数参数：

- 单向读通道
- 单向写通道

**select**

当程序中有多个 channel 协同工作，chan1、chan2，在某一时刻，chan1 或 chan2 被触发了，程序要做响应的处理

使用 select 来监听多个 channel，当某个 channel 被触发时（写入、读出、关闭）进行处理

select 语法与 switch...case 很像，但所有分支条件都必须是通道 io

## 六、网络编程

### 1.网络分层

<table>
    <tr>
        <th>OSI 七层网络模型</th>
        <th>TCP/IP 四层概念模型</th>
        <th>对应网络协议</th>
    </tr>
    <tr>
        <th>应用层（Application）</th>
        <th rowspan="3">应用层</th>
        <th>HTTP、TFTP、NFS、WAIS、SMTP</th>
    </tr>
    <tr>
        <th>表示层（Presentation）</th>
        <th>Telnet、Rlogin、SNMP、Gopher</th>
    </tr>
    <tr>
        <th>会话层（Session）</th>
        <th>SMTP、DNS</th>
    </tr>
    <tr>
        <th>传输层（Transport）</th>
        <th>传输层</th>
        <th>TCP、UDP</th>
    </tr>
    <tr>
        <th>网络层（Network）</th>
        <th>网络层</th>
        <th>IP、ICMP、ARP、RARP、AKP、UUCP</th>
    </tr>
    <tr>
        <th>数据链路层（Data Link）</th>
        <th rowspan="2">数据链路层</th>
        <th>FDDI、Ethernet、Arpanet、PDN、SLIP、PPP</th>
    </tr>
    <tr>
        <th>物理层（Physical）</th>
        <th>IEEE 802.1A，IEEE 802.2~802.11</th>
    </tr>
</table>


### 2.socket

<img src="https://i.loli.net/2021/10/18/F8zkLiAreg2o71J.jpg" alt="preview" style="zoom: 50%;" />

### 3.http

#### ①.概述

底层仍是 TCP，可以理解为一个短链接的 TCP

编写 web 的语言：

- Java
- PHP，现在在使用 go 重构
- Python
- Go ==> beego、gin 两个主流的 web 框架

HTTP 协议：我们使用浏览器访问的时候发送的就是 HTTP 请求

1. HTTP 是应用层的协议，底层还是依赖传输层：TCP（短链接），网络层（IP）
2. 是无状态的，每一次请求都是独立的，下次请求需要重新建立连接
3. HTTPS 不是标准协议，HTTP 是标准协议；HTTPS=HTTP+SSL（非对称加密，数字证书）；现在所有网站都会尽量要求使用 HTTPS 开发：安全

#### ②.http 请求报文格式

![HTTP 请求报文由请求行、请求头部、空行 和 请求包体 4 个部分组成](https://i.loli.net/2021/10/18/TF24Q6sEJKrOe8N.jpg)

一个 HTTP 请求可以分为四部分

1. 请求行：包含三部分

   1. 格式：**请求方法**+空格+**URL**+空格+**协议版本号**+**CRLF**
   2. 请求方法
      1. GET：获取数据
      2. POST：上传数据（表单、JSON）
      3. PUT：修改数据
      4. DELETE：删除数据
      5. ...
   3. URL
   4. 协议版本

2. 请求头

   1. 格式：**key:value**
   2. 可以有很多键值对（包含协议自带的和用户自定义的）
   3. 常见重要字段
      1. Accept：接受数据的格式
      2. User-Agent：用户浏览器引擎
      3. Connection：Keep-Alive（长链接）、Close（短链接）
      4. Accept-Encoding：gzip、...，描述客户端可以接受的编码
      5. Cookie：由服务器设置的 key=value 数据，客户端下次请求的时候可以携带过来
      6. Content-Type：application/-form（表示上传的数据是表单格式）、application/-json（表示 body 中的数据是 JSON 格式）

3. 空行

   告诉服务器请求头结束，用于分割请求头和请求包体

4. 请求包体（可选）

   1. 一般在 POST 时会配套 Body，其他方法也可以但语义不明确，容易混淆
   2. 上传两种数据格式：表单、JSON

**前端与后端传递数据的方法：**

1. 放在请求头中
2. 放在请求体中
3. 放在 URL 中：`GET /user?id=1001&score=90`
   1. ? 分割参数和 URL
   2. 多个参数之间用 & 分割，每一个参数都是一个键值对

#### ③.http 响应报文格式

![img](https://i.loli.net/2021/10/18/O73WxNlAXekymD9.png)

一个 HTTP 响应体包含四部分

1. 状态行

   1. 格式：**协议版本号**+空格+**状态码**+空格+**状态描述**+CRLF
   2. 常用状态码：
      - 1xx：客户端可以继续发送请求
      - 2xx：正常访问，200
      - 3xx：重定向
      - 4xx：
        - 401：未授权 Not Authorized
        - 404：Not Found
      - 5xx：
        - 501：服务器内部错误

2. 响应头

   1. 格式：**key:value**

3. 空行

   分割响应头和响应包体

4. 响应包体（可选）

   1. 通常返回 JSON 数据

### 4.JSON

JSON 语法要求最后一个元素后面不能加逗号

**JSON 的编解码**

在网络传输过程中，把结构体编码成 JSON 字符串传输

接受 JSON 字符串后，需要把字符串转换成结构体然后操作

注意结构体中小写字母开头的字段不会被编码，同样解码也不会存在

### 5.结构体的标签

```go
type Teacher struct {
	Name    string `json:"-"`                 //在使用 JSON 编码时不编码该字段
	Subject string `json:"Sub_name"`          //在 JSON 编码时会用标签替换原字段名
	Age     int    `json:"age,string"`        //在 JSON 编码时更改字段数据类型,注意中间逗号左右不能加空格:字段名,字段类型
	Address string `json:"address,omitempty"` //omitempty表示如果该字段为空则不参与编码
	gender  string
}
```

如果 JSON 格式要求 key 小写或其他名称可以通过标签进行自定

更换数据类型也可以通过标签进行更改

`omitempty` 标签可以指定字段在 JSON 编码时不参与

