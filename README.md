# Golang 学习笔记

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

```go
go build //编译自身包和依赖包
go install //编译并安装自身包和依赖包
gofmt //格式化代码
go doc //生成代码文档
```

## 三、GO 基本语法

### 1.文件名、关键字和标识符

#### 文件名

Go 的源文件以 `.go` 为后缀名存储，文件名均有小写字母组成，允许使用下划线，不包含空格或其他特殊字符，且必须以小写字母开头，eg：`test_1.go`

#### 关键字

|  **break**   |   **default**   |  **func**  | **interface** | **select** |
| :----------: | :-------------: | :--------: | :-----------: | :--------: |
|   **case**   |    **defer**    |   **go**   |    **map**    | **struct** |
|   **chan**   |    **else**     |  **goto**  |  **package**  | **switch** |
|  **const**   | **fallthrough** |   **if**   |   **range**   |  **type**  |
| **continue** |     **for**     | **import** |  **return**   |  **var**   |

#### 标识符

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

