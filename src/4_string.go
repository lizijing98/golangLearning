package main

import "fmt"

func main() {

	//1.定义
	name := "lizijing"
	// 需要换行，原生输出字符串使用反引号``
	usage := `./a.out<option>
	-h help
	-o other`
	fmt.Println("name:", name)
	fmt.Println("usage:", usage)

	//2.长度、访问
	//C++: name.length
	//go: string 没有 length 方法，可以使用自由函数 len() 进行处理
	//一样会有访问下标越界的问题
	length := len(name)
	for i := 0; i < length; i++ {
		fmt.Printf("i: %d, v: %c \n", i, name[i])
	}

	//3.拼接
	i, j := "hello", "go"
	fmt.Println("i+j:", i+j)

	//4.字符串常量
	const address = "Nanjing"
	//address="shanghai" 会报错
	fmt.Println("address:", address)
}
