package main

import (
	"fmt"
	"runtime"
)

func show(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Printf(msg + "\n")
	}
	if i > 5 {
		//如果大于5就退出携程
		runtime.Goexit()
	}
}
func main() {
	//开启子协程
	go show("java\n")

}
