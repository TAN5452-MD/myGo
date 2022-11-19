package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Print("\n")
		fmt.Print(msg)
		time.Sleep(time.Millisecond * 100)
	}
}
func main() {
	//main函数退出 程序就结束了 协程自动死亡
	go showMsg("java") // 通过在前面添加go关键字开启协程
	showMsg("golang")
}
