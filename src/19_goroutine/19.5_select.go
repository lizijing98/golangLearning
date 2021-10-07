package main

import (
	"fmt"
	"time"
)

//当程序中有多个 channel 协同工作，chan1、chan2，在某一时刻，chan1 或 chan2 被触发了，程序要做响应的处理
//使用 select 来监听多个 channel，当某个 channel 被触发时（写入、读出、关闭）进行处理
//select 语法与 switch...case 很像，但所有分支条件都必须是通道 io

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	//1.启动一个 goroutine，负责监听两个 channel
	go func() {
		for true {
			fmt.Println("==============监听中==============")
			select {
			case data1 := <-chan1:
				fmt.Println("1111 从 chan1 读取数据成功，data1:", data1)
			case data2 := <-chan2:
				fmt.Println("2222 从 chan2 读取数据成功，data2:", data2)
			default:
				fmt.Println("==============监听中==============")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	//2.启动 goRou1，写 chan1
	go func() {
		for i := 0; i < 10; i++ {
			chan1 <- i
			time.Sleep(1 * time.Second / 2)
		}
	}()
	//2.启动 goRou2，写 chan2
	go func() {
		for i := 0; i < 10; i++ {
			chan2 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	for true {
	}
}
