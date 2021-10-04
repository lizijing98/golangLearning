package main

import "fmt"

func main() {
	// go 中存在指针，与 C 中指针不同的是不需要去释放（自动释放），不会造成野指针（内存泄露）
	// C 语言中结构体调用时：ptr->name
	// go 语言中： ptr.name
	// go 中使用指针时，会使用内部的垃圾回收机制(gc:garbage collector),开发人员不需要手动释放
	// C 中不允许返回栈上的指针，go 中允许，程序会在编译时确定这个变量的分配位置，编译时发现有必要的情况下就分配到堆上

	// 1.直接定义
	name := "lizijing"
	ptr1 := &name

	fmt.Println("name:", *ptr1)
	fmt.Println("name ptr1:", ptr1)

	// 2.使用 new 关键字定义
	ptr2 := new(string)
	*ptr2 = "LIZIJING"
	fmt.Println("name2:", *ptr2)
	fmt.Println("name2 ptr:", ptr2)

	// 可以返回栈上指针
	res := testPtr()
	fmt.Println("res city:", *res)
	fmt.Println("res city ptr:", res)

	// 判断空指针
	// C: null、C++: nullptr、go:nil
	if res == nil {
		fmt.Println("res 是空，nil")
	} else {
		fmt.Println("res city:", *res)
	}
}

// 返回栈上指针
// 定义一个函数，返回一个 string 类型的指针，go 语言返回写在参数列表后
func testPtr() *string {
	city := "南京"
	cityPtr := &city
	return cityPtr
}
