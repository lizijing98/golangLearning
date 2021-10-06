package classSrc

// go 语言中权限通过首字母大小写控制
//1.import-->如果包名不同，只有大写字母开头的才是 public
//2.对于类里面的成员、方法-->只有大写字母开头的才能在其他包中使用

type Person struct {
	Name   string
	Age    int
	Gender string
	Score  float64
}

//定义一个嵌套 Person 的子类 Student

type Student struct {
	Person Person //包含 Person 类型的变量，此时属于类的嵌套
	School string
}

//定义一个'继承' Person 的子类 Teacher
//注意不是真正的继承，只是 struct 之间的组合

type Teacher struct {
	Person  // '继承'直接写类型，不写字段名
	subject string
}

func (teacher *Teacher) SetSubject(subject string) {
	teacher.subject = subject
}
