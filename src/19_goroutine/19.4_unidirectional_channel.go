package main

import (
	"fmt"
	"time"
)

func main() {
	//numsChan := make(chan int, 10) //双向通道，可读可写

	//单向读通道
	//var ReadChan <-chan int
	//单向写通道
	//var WriteChan chan<- int

	//生产者消费者模型
	//C: 数组+锁 thread1：写 thread2：读
	//go: goroutine+channel

	//1.在主函数中创建一个双向通道 numsChan
	//2.将 numsChan 传递给 producer 负责生产
	//3.将 numsChan 传递给 consumer 负责消费

	numsChan := make(chan int, 10)
	//双向通道可以赋值给同类型的单向通道，反之不行
	go producer(numsChan)
	go consumer(numsChan)

	time.Sleep(1 * time.Second)
}

//producer 生产者 ===> 提供一个只写的通道
func producer(output chan<- int) { //数据单向写到通道中
	for i := 0; i < 20; i++ {
		output <- i
		//date := <-output // 单向写通道不允许有读操作
		fmt.Println("==>向管道中写入数据:", i)
	}
}

//consumer 消费者 ===> 提供一个只读的通道
func consumer(input <-chan int) {
	for value := range input {
		fmt.Println("<==从管道中读出数据:", value)
	}
}
