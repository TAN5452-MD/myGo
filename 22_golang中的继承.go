package main

import "fmt"

type Person struct {
	//如果想要变成公共的 首字母大写
	name string
	age  int
}
type Student struct {
	p     Person //可以理解为继承
	color string
}

func main() {
	/*stu := Student{
		p:     Person{name: "tom", age: 18},
		color: "blue",
	}*/
	//顺序初始化
	stu := Student{
		Person{name: "tom", age: 18},
		"blue",
	}
	fmt.Print(stu.p.name)
}
