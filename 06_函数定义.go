package main

import "fmt"

func main() {
	var s int = sum(100, 20000000000000)
	fmt.Printf("%v", s)
}
func sum(a int, b int) int {
	return a + b
}
func notRes() (a int, b int) {
	//等价于sum
	a = 100
	b = 200
	return
}

// 剩余参数
func rest(args ...int) {
	for i, arg := range args {
		fmt.Print(arg)
	}
}
