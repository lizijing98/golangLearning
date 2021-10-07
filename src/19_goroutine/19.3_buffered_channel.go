package main

import (
	"fmt"
	"time"
)

func main() {
	//numsChan := make(chan int, 10)
	//1.当缓冲区写满时，写阻塞，当被读取后，再恢复写入
	//2.当缓冲区读取完毕，读阻塞
	//3.如果管道没有使用 make 分配空间，那么管道默认是 nil 的，读取写入都会阻塞
	//4.对于一个 channel，读与写次数必须对等，否则会造成死锁

	var names chan string //没有分配空间的 channel 无法写入
	names = make(chan string, 10)

	go func() {
		fmt.Println("names:", <-names)
	}()

	names <- "hello" //如果 channel 是 nil 的，写操作会阻塞在这里
	time.Sleep(1 * time.Second)

	numsChan := make(chan int, 10)

	//写 10 次
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("==>子 go 程写入:", i)
			numsChan <- i
		}
		fmt.Println("写入完毕，关闭管道")
		close(numsChan)
	}()

	//读 20 次
	//如果在一个 go 程中进行 channel 不对称读写，不会造成死锁，但会内存泄露
	//使用 for:range 解决读写不一致
	func() {
		/*for i := 0; i < 20; i++ {
			fmt.Println("<==主 go 程读出:", <-numsChan)//此时只能读到 10 次，再往后读会造成死锁
		}*/
		// for:range 不知道 channel 是否写完，会一直等待
		//解决：写入端在写数据完毕时，进行管道关闭，for:range 遍历关闭的管道后，会退出
		for value := range numsChan {
			fmt.Println("<==主 go 程读出:", value)
		}
	}()

	time.Sleep(1 * time.Second)
}
