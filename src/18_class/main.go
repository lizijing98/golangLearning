package main

import (
	"18_class/classSrc"
	"fmt"
)

func main() {
	stu1 := classSrc.Student{
		Person: classSrc.Person{
			Name:   "小李",
			Age:    23,
			Gender: "男",
			Score:  100,
		},
		School: "NUIST",
	}

	tea1 := classSrc.Teacher{
		Person: classSrc.Person{
			Name:   "老李",
			Age:    40,
			Gender: "男",
		},
	}
	tea1.SetSubject("Golang")

	fmt.Println("stu1:", stu1)
	fmt.Println("stu1 Name:", stu1.Person.Name)
	fmt.Println("tea1:", tea1)
	fmt.Println("tea1 Name:", tea1.Name)

	//'继承'的时候虽然没有定义字段名，但是还是会自动创建一个同名字段
	//在子类中依然可以操作父类
	fmt.Println("tea1 person:", tea1.Person)

	James := classSrc.Person{
		Name:   "James",
		Gender: "男",
	}
	fmt.Println("参数使用对象：")
	James.Eat()                         //LiLi is eating
	fmt.Println("Name is:", James.Name) //Name is: James
	fmt.Println("参数使用指针：")
	James.EatPointer()                  //LiLi is eating
	fmt.Println("Name is:", James.Name) //Name is: LiLi
}
