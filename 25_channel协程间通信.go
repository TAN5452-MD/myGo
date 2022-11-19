package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel通道就是让协程互相通信

// 创建int类型通道，只能传入int类型
var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 5)
	values <- value
}
func main() {
	//从通道接受值
	defer close(values) //关闭通道
	go send()
	fmt.Print("wait\n")
	value := <-values
	fmt.Printf("%v\n", value)
	fmt.Print("end")
}
