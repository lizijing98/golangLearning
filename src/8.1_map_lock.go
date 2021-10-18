package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 7.并发处理时，需要对 Map 进行上锁
	// map 不允许同时进行读写，存在多个 goroutine 同时操作的情况，需要对 map 进行上锁
	var lock sync.RWMutex
	var names = make(map[int]string)

	go func() {
		fmt.Println("goroutine 1")
		for true {
			lock.Lock()
			names[0] = "xiaoming"
			lock.Unlock()
		}
	}()

	go func() {
		fmt.Println("goroutine 2")
		for true {
			lock.Lock()
			fmt.Println("names:", names[0])
			lock.Unlock()
		}
	}()

	for true {
		fmt.Println("over")
		time.Sleep(1 * time.Second)
	}

}
