package main

import (
	"fmt"
	"time"
)

func main() {
	//实例化一个ticker
	ticker := time.NewTicker(time.Second)
	counter := 1
	//ticker 就等于setInterval
	for _ = range ticker.C {
		fmt.Println("ticker")
		counter++
		if counter > 5 {
			ticker.Stop() //停止ticker
			break
		}
	}

	go foo()
	sum := 0
	for v := range chanInt {
		fmt.Println("收到了", v)
		sum += v
		if sum > 10 {
			break
		}
	}
}

var chanInt = make(chan int)

func foo() {
	ticker := time.NewTicker(time.Second)
	//一定是.C
	for _ = range ticker.C {
		select {
		case chanInt <- 1:
		case chanInt <- 2:
		case chanInt <- 3:
		}
	}
}
