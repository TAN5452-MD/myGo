package main

//waitGroup实现同步
import (
	"fmt"
	"sync"
)

// 定义一个waitGroup
var wg sync.WaitGroup

func showMsg(i int) {
	//完成就调用done方法 协程-1
	defer wg.Done()
	fmt.Printf("%v\n", i)
}
func main() {
	for i := 0; i < 10; i++ {
		//开启一个协程就 wg+1
		wg.Add(1)
		go showMsg(i)
	}
	//等带waitGroup中的协程结束
	wg.Wait()
	fmt.Printf("end")

}
