package main

import "fmt"

//实现 go 多态需要实现定义接口
//eg:人类的武器发起攻击，不同等级子弹效果不同

//定义一个接口，注意类型是 interface

type IAttack interface {
	Attack()
	//接口的函数可以有多个，但只能有函数原型，不能有实现
	//如果定义了多个接口函数，则需要实现接口的类需要实现所有函数，才可以赋值
}

type HumanLowLevel struct {
	name  string
	level int
}

func (h *HumanLowLevel) Attack() {
	fmt.Println("我是:", h.name)
	fmt.Println("等级为:", h.level)
}

type HumanHighLevel struct {
	name  string
	level int
}

func (h *HumanHighLevel) Attack() {
	fmt.Println("我是:", h.name)
	fmt.Println("等级为:", h.level)
}

// DoAttack 定义一个多态的接口，传入不同的对象，调用同名方法，实现不同效果
func DoAttack(a IAttack) {
	a.Attack()
}

func main() {
	lowLevel := HumanLowLevel{
		name:  "player1",
		level: 10,
	}
	lowLevel.Attack()
	highLevel := HumanHighLevel{
		name:  "player2",
		level: 100,
	}

	// 对 player 赋值为 lowLevel,接口需要指针来赋值？
	var player IAttack
	player = &lowLevel
	player.Attack()
	player = &highLevel
	player.Attack()
	fmt.Println("多态...")
	DoAttack(&lowLevel)
	DoAttack(&highLevel)
}
