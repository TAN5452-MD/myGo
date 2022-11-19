package main

import "fmt"

// 函数类型和函数变量
func main() {
	//定义
	type f1 func(a int, b int) int
	var instanceF1 f1
	//重载
	instanceF1 = sum
	r := instanceF1(1, 2)
	fmt.Print(r)

	instanceF1 = max
	r2 := instanceF1(1, 2)
	fmt.Print(r2)
}
func sum(a int, b int) int {
	return a + b
}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
