package classSrc

import "fmt"

//定义一个 Person 类，包含方法：Eat、Run、Laugh、成员
// public private
/*C++:
class Person{
	private ... ...;
	public ... {}
	private ... {}
}*/

/*type Person struct {
	name   string
	age    int
	gender string
	score  float64
}*/

// go 中对象的方法不在结构体中写，在结构体外写再绑定

func (p *Person) EatPointer() {
	// 类的方法可以使用自己的成员/指针
	// 使用指针操作对象本身，不使用指针操作对象的拷贝
	p.Name = "LiLi"
	fmt.Printf("%s is eating\n", p.Name)
}
func (p Person) Eat() {
	// 类的方法可以使用自己的成员/指针
	// 使用指针操作对象本身，不使用指针操作对象的拷贝
	p.Name = "LiLi"
	fmt.Printf("%s is eating\n", p.Name)
}
