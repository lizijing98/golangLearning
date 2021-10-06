package main

import "fmt"

// C 语言中可以使用 typedef 定义类型：typedef Int MyInt

type MyInt int // go 中定义类型

// C 中使用结构体 struct 来模拟类
/*C:
struct Person{
	...
}*/
// go 中使用 type+struct 来处理结构体

type Person struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {
	var i, j MyInt
	i, j = 10, 20
	fmt.Println("i:", i, "j:", j)

	// 结构体赋值时每个字段都赋值，则字段名可以省略不写
	xiaoming := Person{
		"xiaoming",
		23,
		"男",
		60, //最后一个元素必须加逗号，如果不加必须与}同行
	}
	// 结构体没有->操作
	fmt.Println("xiaoming:", xiaoming)
	fmt.Println("xiaoming name:", xiaoming.name)

	s1 := &xiaoming
	fmt.Println("使用指针打印：")
	fmt.Println("xiaoming age:", s1.age)
	fmt.Println("xiaoming gender:", (*s1).gender)

	// 对部分字段赋值则必须加上指定的 key
	xiaowang := Person{
		name:   "xiaowang",
		age:    22,
		gender: "男生"}
	fmt.Println("xiaowang:", xiaowang)
	xiaowang.score = 90
	fmt.Println("xiaowang:", xiaowang)
}
