package main

import (
	"fmt"
	"sync"
	"time"
)

var wd sync.WaitGroup

// 定义互斥锁
var Lock sync.Mutex
var i int = 100

func add() {
	//协程完成
	defer wd.Done()
	//当一个协程进入就加锁防止其他协程操作
	Lock.Lock()
	i++
	//操作完成就解锁
	Lock.Unlock()
}
func sub() {
	//协程完成
	defer wd.Done()
	Lock.Lock()
	i--
	Lock.Unlock()

}
func main() {
	for i := 0; i < 100; i++ {
		go add()
		//协程+1
		wd.Add(1)
		go sub()
		wd.Add(1)
	}
	time.Sleep(time.Millisecond * 10)
	//等待协程执行完成
	wd.Wait()
	fmt.Printf("i=: %v", i)

}
