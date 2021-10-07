package main

import "fmt"

func main() {
	// 定义接口
	var i, j, k interface{}
	names := []string{"小李", "小王"}
	i = names
	fmt.Println("i 此时代表切片数组:", i)
	age := 20
	j = age
	fmt.Println("j 此时代表 int 数字:", j)
	str := "hello"
	k = str
	fmt.Println("k 此时代表字符串:", k)

	//现在只知道 k 是 interface 不知道 k 代表的数据的类型
	//快速判断 k 是否为某个数据类型
	kvalue, isInt := k.(int)
	if isInt {
		fmt.Println("k 代表的数据的类型是 int:", kvalue)
	} else {
		fmt.Println("k 代表的数据的类型不是 int:", kvalue)
	}

	//最常用的场景：把 interface 当做一个函数的参数，类似于 print
	//使用 switch 来判断用户输入的不同类型，根据不同的数据类型做相应的处理
	//创建一个具有三个接口类型的切片
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "Hello go"
	array[2] = true

	for _, value := range array {
		switch value.(type) { //.(type)获取实际数据类型
		case int:
			fmt.Printf("当前数据类型为 int:%d\n", value)
		case string:
			fmt.Printf("当前数据类型为 string:%s\n", value)
		case bool:
			fmt.Printf("当前数据类型为 bool:%v\n", value)
		case float64:
			fmt.Printf("当前数据类型为 float:%f\n", value)
		default:
			fmt.Println("未处理的数据类型")
		}
	}
}
