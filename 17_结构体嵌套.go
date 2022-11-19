package main

import "fmt"

type Walk struct {
	go1 string
}
type Dog struct {
	name string
	walk Walk
}

func main() {
	//要先实现子结构体
	walk := Walk{go1: "run"}
	dog := Dog{name: "wangwang", walk: walk}
	fmt.Print(dog.walk)
}
