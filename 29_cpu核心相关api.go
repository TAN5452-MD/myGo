package main

import (
	"fmt"
	"runtime"
)

func main() {
	//查看cpu核心数
	fmt.Printf("CPU核心数：%v", runtime.NumCPU())
	//设置最大cpu核心数
	runtime.GOMAXPROCS(1)
}
