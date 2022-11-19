package main

import (
	"fmt"
	"runtime"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf(msg + "\n")
	}
}
func main() {
	//开启子协程
	go show("java\n")
	for i := 0; i < 2; i++ {
		//通过api让出cpu时间片给子协程使用
		runtime.Gosched()
		//fmt.Print("goland\n")
	}
	fmt.Print("end\n")
}
