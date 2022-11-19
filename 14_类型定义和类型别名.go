package main

import "fmt"

func main() {
	//类型定义
	type myInt int
	var a myInt = 1
	fmt.Printf("%T", a)
	//类型别名
	type Myint = int
}
