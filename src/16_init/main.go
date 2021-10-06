package main

import (
	ADD "16_init/add"
	_ "16_init/sub" //使用_只会隐式调用 sub 包中的 init 函数，编译不会报错
	//.  "11_import/sub" //. 代表用户在调用这个包中的函数中，不需要使用包名.的形式，不建议使用，类似于 using namespace
	"fmt"
) // 多个 import 可以用()进行包裹

func main() {
	// 包中的函数名一定要大写字母开头，否则无法被调用
	// 如果一个包中的函数想对外提供访问，则首字母一定要大写，类似于 public
	// 小写字母开头的函数类似于 private，只有相同包名的文件才能使用
	//subNum := SUB.Sub(1, 5)
	addNum := ADD.Add(2, 6)
	//fmt.Printf("subNum:%d\n", subNum)
	fmt.Printf("addNum:%d\n", addNum)
}
