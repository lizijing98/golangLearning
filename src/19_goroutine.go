package main

import (
	"fmt"
	"time"
)

//这个将用于子 go 程
func display(i int) {
	count := 1
	for {
		fmt.Println("===>这是子 go 程，i=", i, "，count:", count)
		count++
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//启动子 go 程
	for i := 0; i < 3; i++ {
		//启动多个子 go 程，他们会公平竞争 CPU 资源
		go display(i)
	}
	/*go func(){
		count := 1
		for {
			fmt.Println("===>这是子 go 程，count:", count)
			count++
			time.Sleep(1 * time.Second)
		}
	}()*/
	//主 go 程
	count := 1
	for {
		fmt.Println("这是主 go 程，count:", count)
		count++
		time.Sleep(3 * time.Second)
	}
}
