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

	var j int64 = 10
	//读操作
	atomic.LoadInt64(&j)
	//写操作
	atomic.StoreInt64(&j)

	//比较并且交换
	res := atomic.CompareAndSwapInt64(&j, 10, 20)
	fmt.Appendln(res)
}

func add1() {
	//cas 比较并且交换，系统会比较old值和new值 成功才会调用
	atomic.AddInt32(&ccount, 1)
}

func sub2() {
	//cas 比较并且交换，系统会比较old值和new值 成功才会才调用
	atomic.AddInt32(&ccount, -1)
}
