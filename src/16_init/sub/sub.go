package sub

import "fmt"

//package utils // go 中同一个目录层级下不允许出现多个包名

// init函数没有参数，没有返回值，原型固定
// 类似构造函数？
// init() 不允许显式调用
func init() {
	fmt.Println("this is the first init() in sub package")
}

func init() {
	fmt.Println("this is a init() in sub package")
}

// 一个包中包含多个 init() 时，都会被调用
func init() {
	fmt.Println("this is another init() in sub package")
}

func Sub(a, b int) int {
	//test_1() //由于 test_1() 与 Sub() 在同一个包下面，所以可以直接使用，且不需要 sub.
	return a - b
}
