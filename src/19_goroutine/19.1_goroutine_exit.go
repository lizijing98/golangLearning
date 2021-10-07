package main

import (
	"fmt"
	"runtime"
	"time"
)

//GOEXIT ==> 提前退出当前 go 程
//return ==> 返回当前函数
//exit ==> 退出当前进程

func main() {
	go func() {
		fmt.Println("这是子 go 程 1")
		go func() {
			fmt.Println("这是子 go 程 2")
			func() {
				// go 程内部的匿名函数，func1
				fmt.Println("这是子 go 程内部的匿名函数")
				//return // 这个 return 结束当前 go 程内部的匿名函数 func1，会打印“子 go 程1/2结束”
				//os.Exit(-1)// 这个 os.Exit 结束当前进程，不会打印“子 go 程1/2结束”和“主 go 程结束”
				runtime.Goexit() // 这个 runtime.Goexit() 结束当前 go 程，不会打印“子 go 程 2 结束”，会打印“子 go 程 1 结束/主 go 程结束”
			}()
			fmt.Println("子 go 程 2 结束")
		}()
		time.Sleep(2 * time.Second)
		fmt.Println("子 go 程 1 结束")
	}()
	fmt.Println("这是主 go 程")
	time.Sleep(2 * time.Second)
	fmt.Println("主 go 程结束")
}
