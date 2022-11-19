package main

import "fmt"

func main() {
	//切片
	//切片的声明 不固定长度
	//var slice1 []string

	var s2 = make([]int, 5)
	fmt.Print(len(s2))
	fmt.Print(cap(s2))
	fmt.Print(s2[0])
}
