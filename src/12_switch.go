package main

import (
	"fmt"
	"os"
)

// 从命令行输入参数并用 switch 进行判断
// C:argc,**argv
// go:os.Args ==> 直接从命令行获取输入，是一个 slice

func main() {
	cmds := os.Args
	// os.Args[0]是程序名，后面是输入的参数

	for index, cmd := range cmds {
		fmt.Printf("index:%d,cmd:%s,cmds len:%d \n", index, cmd, len(cmds))
	}

	if len(cmds) < 2 {
		fmt.Println("要求输出参数！")
	}

	// go 中默认 switch 已经带上 break，不需要手动处理
	// 想要向下穿透需要加上关键字: fallshrough
	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		fallthrough
	case "go":
		fmt.Println("go")
	case "1234":
		fmt.Println("1234")
	default:
		fmt.Println("cmds:", cmds)
	}
}
