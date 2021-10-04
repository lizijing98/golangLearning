package main

import "fmt"

func main() {
	// 不定长数组也称为切片: slice，本质上也是数组，可以动态改变长度
	// 1.定义，eg:定义一个包含多个地名的 slice
	// addsFixed := [10]string{"Beijing", "Shanghai", "Guangzhou"} // 定长
	addsVariable := []string{"Chengdu", "Nanjing"}

	for index, value := range addsVariable {
		fmt.Println("add[", index, "]", value)
	}

	// 2.追加
	fmt.Println("追加")
	// addsVariable = append(addsVariable, "Hangzhou") // 向自身追加一个元素
	addsVariable1 := append(addsVariable, "Hangzhou") // 追加一个元素并赋值给一个新变量
	fmt.Println("addsVariable:", addsVariable)
	fmt.Println("addsVariable1:", addsVariable1)

	// 3.对于一个切片，不仅有长度的概念(len())，还有容量的概念(cap())

	fmt.Println("容量")
	fmt.Println("追加元素前:")
	fmt.Println("length:", len(addsVariable))
	fmt.Println("cap:", cap(addsVariable))
	fmt.Println("追加一个元素后:")
	fmt.Println("length:", len(addsVariable1))
	fmt.Println("cap:", cap(addsVariable1))
	// go 为了保障效率直接分配多一倍的容量，但长度还是实际的使用长度，在容量为达到当前上限前不会再分配，直至达到上限
	// 1k 以下 *2,1k 以上 *2 or *>1.25
	fmt.Println("追加两个元素后:")
	addsVariable1 = append(addsVariable1, "Zhengzhou")
	fmt.Println("length:", len(addsVariable1))
	fmt.Println("cap:", cap(addsVariable1))
	fmt.Println("追加三个元素后:")
	addsVariable1 = append(addsVariable1, "Beijing")
	fmt.Println("length:", len(addsVariable1))
	fmt.Println("cap:", cap(addsVariable1))

	var nums []int
	for i := 0; i < 50; i++ {
		nums = append(nums, i)
		fmt.Println("lens:", len(nums), "cap:", cap(nums))
	}
}
