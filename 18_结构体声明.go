package main

import "fmt"

type Person struct {
	name string
}

// per Person 接收者 receiver 声明方法所属的结构体
// 并非要属于同一个文件，但一定要属于同一个包
func (per Person) eat() {
	fmt.Print("eat....")
}

func main() {
	per := Person{
		name: "tom",
	}
	per.eat()
}
