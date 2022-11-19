package main

import "fmt"

// 高阶函数
func main() {
	hello("tom", sayH)
	r := closure()
	r(200)
}
func sayH(name string) {
	fmt.Print(name)
}
func hello(name string, f func(string)) {
	f(name)
}

func closure() func(int) int {
	a := 200
	return func(y int) int {
		a = a + y
		fmt.Print("\n")
		fmt.Print(a)
		return a
	}
}
