package main

import "fmt"

//go 中的函数可以返回多个返回值，返回值在参数列表之后
//如果有多个返回值，需要用()进行包裹，参数、返回值之间用,分割
//如果只有一个返回值，且没有名字，则不需要加括号
func test1(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}

//相同类型的参数可以省略部分参数声明
//存在返回值名称，可以直接参与运算
func test2(a, b int, c string) (num2 int, str2 string) {
	num2 = a / b
	str2 = c
	return //存在返回值名，可以直接写 return
}

func main() {
	num1, str1, _ := test1(1, 2, "hello")
	fmt.Printf("test1:num1:%d,str1:%s\n", num1, str1)
	num2, str2 := test2(1, 2, "test2:")
	fmt.Printf("%s%d\n", str2, num2)
}
