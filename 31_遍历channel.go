package main

import "fmt"

var c = make(chan int)

func main() {
	//如果没有关闭通道，超过读写数量时，会报死锁错误
	//如果没有关闭通道，超过读写数量时，赋给默认值
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		v, ok := <-c
		if ok {
			fmt.Print(v)
		}
	}
	/*for v := range c {
		fmt.Print(v)
	}*/
	/*	for i := 0; i < 3; i++ {
		r := <-c
		fmt.Print(r)
	}*/
}
