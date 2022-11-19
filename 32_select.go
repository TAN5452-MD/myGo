package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanString = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanString <- "hell0"
	}()
	for {
		//select语句是随机的
		select {
		case r := <-chanInt:
			fmt.Print(r)
		case r := <-chanString:
			fmt.Print(r)
		default:
			fmt.Print("defalut")

		}
		time.Sleep(time.Second)
	}

}
