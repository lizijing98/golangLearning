package main

import (
	"fmt"
	"os"
)

func main() {
	readFile("hello.go")
}

func readFile(filename string) {
	//1.go 语言一般将错误码作为第二个返回值
	//2.当 err 为 nil 时代表没有错误，执行成功，非 nil 表示执行失败
	file1, err1 := os.Open(filename)

	//一个函数没有函数名就是匿名函数，加上()表示调用，一次性的函数，可以传参
	//类似 lambda 表达式？
	defer func(str string) {
		file1.Close()
		fmt.Printf("close file: %s", str)
	}("000") // 创建一个匿名函数同时调用

	defer fmt.Println("111")
	defer fmt.Println("222")
	// 注意 defer 的执行顺序：222->111->close file
	// 类似于栈的执行顺序：先进后出

	if err1 != nil {
		fmt.Println("os.Open(\"hello.go\") failed:\n", err1)
		return
	}
	buf := make([]byte, 1024)
	n, _ := file1.Read(buf)
	fmt.Println("读取文件的实际长度:", n)
	fmt.Println("读取文件的内容:\n", string(buf))
}
