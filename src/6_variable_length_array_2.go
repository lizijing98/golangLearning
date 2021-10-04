package main

import "fmt"

func main() {
	adds := [6]string{"北京", "上海", "广东", "深圳", "杭州", "南京"}

	// 基于 adds 创建一个新的数组

	// 切片可以基于一个数组灵活地创建新的数组，是 adds 的一部分，修改 name1，adds 也会发生改变
	// 类似于引用/浅拷贝
	// 从 0 元素开始截取，:左边的数字可以省略；若一直取到最后元素，:右边的数字可以省略；若全部取用，则:左右两边都可以省略
	name1 := adds[0:2] //左闭右开，左边要右边不要，即取了 adds[0],adds[1]，共两个
	fmt.Println("name1:", name1)
	name2 := adds[1:3] //adds[1],adds[2]
	fmt.Println("name2", name2)

	//修改 name1
	name1[1] = "hello"
	fmt.Println("修改 name1[1]:")
	fmt.Println("adds:", adds)
	fmt.Println("name1:", name1)
	fmt.Println("name2", name2)

	// 字符串也可以进行处理
	str1 := "hello_go_world!"
	str2 := str1[:8]
	fmt.Println("str2:", str2)
	str2 = str1[6:]
	fmt.Println("str2:", str2)

	// 可以创建空切片的时候，明确指定切片容量，减少资源开销，提高运行效率
	// 动态分配造成资源消耗
	name3 := make([]string, 4, 8) // 数据类型、长度、容量，容量不是必须的参数，默认值与长度相同
	fmt.Println("name3:", name3, "length:", len(name3), "cap:", cap(name3))

	// 想让 slice 完全独立于原数组使用 copy()，实现深拷贝
	name4 := make([]string, len(adds))
	copy(name4, adds[:]) // 主要copy()参数接收切片，需要[:]将数组变为切片
	fmt.Println("name4:", name4)
	fmt.Println("改变 name4 中的值")
	name4[2] = "GO"
	fmt.Println("adds:", adds)
	fmt.Println("name4:", name4)
}
