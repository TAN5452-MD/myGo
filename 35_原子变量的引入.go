package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var ccount int32 = 100

func main() {
	for i := 0; i < 100; i++ {
		go add1()
		go sub2()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(ccount)
}

func add1() {
	//cas 比较并且交换，系统会比较old值和new值 成功才会调用
	//加
	atomic.AddInt32(&ccount, 1)
}

func sub2() {
	//cas 比较并且交换，系统会比较old值和new值 成功才会才调用
	atomic.AddInt32(&ccount, -1)
}
