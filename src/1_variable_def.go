package main

import "fmt"

func main() {
	// 变量定义:var 常量定义:const
	// 1.先定义变量再赋值
	var name string
	name = "lizijing"

	fmt.Println("name:", name)

	// 2.定义时直接赋值
	var age = 23 // 可以省略类型
	fmt.Printf("age: %d\n", age)

	// 3.定义直接复制，使用自动推导
	gender := "男"
	fmt.Println("gender:", gender)

	// 4.平行赋值
	i, j := 10, 20 // 同时赋值两个变量
	fmt.Printf("变换前:i: %d,j: %d\n", i, j)
	i, j = j, i
	fmt.Println("变换后:i:", i, "j:", j)
}
