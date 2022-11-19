package main

import "fmt"

// 变量初始化-> init > main
func main() {
	/*	i := 10
		var p *int = &i
		fmt.Print(p)*/
	arrPointer()
}

func arrPointer() {
	var arr = [5]int{1, 2, 3, 4, 5}
	//定义数组指针，go和c++的数组指针不同，c++是指向首地址 而go是指向每个元素
	//go中的指针不允许参与运算
	var ptr [5]*int
	for i := 0; i < len(arr); i++ {
		ptr[i] = &arr[i]
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(*ptr[i])
	}
}
