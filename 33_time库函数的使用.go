package main

import (
	"fmt"
	"time"
)

func main() {
	//实例化一个time
	/*	timer := time.NewTimer(time.Second * 2)
		fmt.Printf("%v\n", time.Now())
		t1 := <-timer.C //阻塞时间,使其完全同步
		fmt.Print(t1)*/

	//<-time.After(time.Second * 2) //time.after的返回值是一个chan time
	//fmt.Print(" ...")

	timer := time.NewTimer(time.Second * 2)
	go func() {
		<-timer.C
		fmt.Print("func")
	}()
	/*	s := timer.Stop()
		if s {
			fmt.Print("stop")
		}*/
	timer.Reset(time.Second * 5)
	time.Sleep(time.Second * 6)
	fmt.Print("done")

}
