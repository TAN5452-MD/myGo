package main

import "fmt"

func main() {
	//最后的defer 最先执行
	fmt.Print(1)
	fmt.Print(2)
	defer fmt.Print(3)
	defer fmt.Print(4)
	fmt.Print(6)

}
