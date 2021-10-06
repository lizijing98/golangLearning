package main

import "fmt"

// go 语言中没有枚举类型，但是可以使用 const+iota（常量累加器）来进行模拟

func main() {
	//var number int
	//var name string
	//var flag bool

	// 使用变量组来统一定义变量
	/*var(
		number int
		name string
		flag bool
	)*/

	// const 常量属于预编译期赋值，不需要:=进行自动推导
	// 使用常量组模拟表示 enum
	const (
		MONDAY    = iota // 0
		TUESDAY          // 1 没有赋值默认与上一行表达式相同
		WEDNESDAY        // 2
		THURSDAY         // 3
		FRIDAY           // 4
		SATURDAY         // 5
		SUNDAY           // 6
		MYDAY     = 8    // 单独赋值则不会与上一行相同
		YOURDAY
	)

	const (
		ZERO = iota     //0
		ONE             //1
		TWO  = iota + 1 //iota=2,TWO=2+1=3
	)

	fmt.Println("MONDAY:", MONDAY, "ONE:", ONE)
	fmt.Println("TWO:", TWO)
	fmt.Println("SUNDAY:", SUNDAY)
	fmt.Println("MYDAY:", MYDAY)
	fmt.Println("YOURDAY:", YOURDAY)
}
