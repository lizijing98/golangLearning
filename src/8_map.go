package main

import "fmt"

func main() {
	// 1.定义字典
	// id <--> names
	//var idNames map[int]string //定义一个 map，此时不能直接使用，它是空的

	// 2.使用 make 分配空间，map 可以不指定长度，建议直接指定长度，节省资源
	//idNames = make(map[int]string, 4)

	// 3.定义时直接分配空间(最常用)
	idNames := make(map[int]string, 4)
	idScores := make(map[int]float64, 4)

	// 4.插入 Map 元素
	// 未分配空间时会报错：panic: assignment to entry in nil map
	idNames[0] = "小李"
	idNames[1] = "小王"

	for index, value := range idNames {
		fmt.Printf("idNames[%d]:%s\n", index, value)
	}

	// 5.确定一个 key 是否存在于 Map 中
	// 在 Map 中不存在访问越界的问题，它认为所有的 Key 都是有效的，所以访问一个不存在的 key 不会崩溃，返回这个类型的零值
	// boolean:false;数字:0;string:空
	fmt.Println("idNames[3]:", idNames[3])    // 访问不存在的 key 值返回零值
	fmt.Println("idScores[10]", idScores[10]) // 访问不存在的 key 值返回零值

	// 存在风险:不能通过获取 value 是否为零值，判断 key 是否存在，可能存在此 key 对应的 value 就是空/零值
	// 需要一个判断 key 是否存在的机制:
	value, isExisted := idScores[4]                              // 如果 key=4 是存在的，则 value 返回 key=4 对应的 value，isExisted 返回 TRUE，否则返回零值和 false
	fmt.Printf("idSocre[4]:%f,isExisted:%t\n", value, isExisted) // %t 布尔类型占位符
	idScores[0] = 86.8
	value, isExisted = idScores[0]
	fmt.Printf("idSocre[0]:%f,isExisted:%t\n", value, isExisted)

	// 6.删除 Map 元素
	// 通过自由函数 delete() 删除
	fmt.Printf("删除前:%s\n", idNames[1])
	delete(idNames, 1)
	fmt.Printf("删除后:%s\n", idNames[1])
	// 删除一个不存在的 key,不会报错
	fmt.Printf("删除无效的 key:%s\n", idNames[5])

	// 7.并发处理时，需要对 Map 进行上锁//TODO
}
