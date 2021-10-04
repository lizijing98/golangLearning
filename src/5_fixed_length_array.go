package main

import "fmt"

func main() {
	//1.定义,eg:定义一个具有十个 int 型数字的定长数组
	//C: int nums[]={1,2,3,4}
	//go:
	// nums := [10]int{1, 2, 3, 4}
	// var nums:=[10]int{1,2,3,4}
	// var nums [10]int =[10]int{1,2,3,4}
	nums := [10]int{1, 2, 3, 4}

	//2.遍历
	//传统方式
	for i := 0; i < len(nums); i++ {
		fmt.Println("num[", i, "]", nums[i])
	}
	//for range
	// index:数组下标
	// value:数组值
	for index, value := range nums {
		// key=0,value=1 -- nums[0]
		// 注意此处 value 为 nums[0] 的一个副本，改变 value 的值 num[0] 的值不会发生改变
		// value 相当于全程只是一个临时变量，再循环过程中被不断赋值
		fmt.Println("value++前,index:", index, "value:", value, "nums:", nums[index])
		value++
		fmt.Println("value++后,index:", index, "value:", value, "nums:", nums[index])
		// go 语言中想忽略某个值可使用下划线(_)代替
		// for _, value := range nums{}
		// 注意 := 两边要存在新变量，如果两个都忽略不能使用 :=,而应该使用 =
		// for _, _ = range nums{}
	}
}
