package main

import "fmt"

func main() {
	/*go 中自增语法只有：i++、i--
	且自增语法必须单独一行*/

	i := 10
	fmt.Println("i:", i) // 输出 i: 10
	i++                  // 自增语句不能和其他语句放在一起
	fmt.Println("i:", i) // 输出 i: 11
}
