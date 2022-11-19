package main

import "fmt"

// 高阶函数
func main() {
	//匿名函数自调用
	func(a int, b int) int {
		fmt.Print(a, b)
		return 0
	}(1, 2)
}
