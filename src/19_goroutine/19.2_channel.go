package main

import (
	"fmt"
	"time"
)

//涉及到多 go 程时，C 语言使用互斥量，上锁来保持资源同步，避免资源竞争问题
//go 语言也支持这种方式，但 go 语言使用更好的方案来解决：管道（channel）
//使用 channel 不需要我们进行加解锁
//A 往 channel 中写数据，B 从 channel 中读数据，go 自动帮我们做好了数据同步

func main() {
	//创建一个装不同类型数据的 channel
	//同 map 一样，使用 channel 时一定要 make，否则结果是 nil
	//不创建空间则是一个无缓冲的 channel
	//有缓冲的 channel 可以实现批量读写
	numChan := make(chan int, 10)
	//strChan := make(chan string)

	//创建三个 go 程，父亲写数据，儿子 1 读数据，儿子 2 写数据
	go func() {
		for i := 0; i < 30; i++ {
			data := <-numChan
			fmt.Println("这是子 go 程 1 读到的数据:", data)
		}
	}()

	go func() {
		for i := 100; i < 110; i++ {
			fmt.Println("这是子 go 程 2:", i)
			numChan <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("这是主 go 程:", i)
		//向管道中写入数据 i
		numChan <- i
	}

	time.Sleep(2 * time.Second)
}
