package main

import "fmt"

// 变量初始化-> init > main
func main() {
	fmt.Print(2)
}
func init() {
	fmt.Print(1)
}

var i int = initVar()

func initVar() int {
	fmt.Print(3)
	return 100
}
