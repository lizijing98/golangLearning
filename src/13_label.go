package main

import "fmt"

func main() {
	// 标签 LABEL1
	// goto LABEL1
	// break LABEL1
	// continue LABEL1

LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//goto LABEL1 // 跳至 LABEL1 时清除上次状态 i=0,j=0
				//continue LABEL1 // 跳至 LABEL1 时记录上次状态 i=1,j=0
				break LABEL1 // 加上 LABEL1 跳出所有循环
			}
			fmt.Printf("i:%d,j:%d\n", i, j)
		}
	}
	fmt.Println("over")
}
