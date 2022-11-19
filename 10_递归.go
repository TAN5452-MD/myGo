package main

import "fmt"

func main() {
	var res int = f2(9)
	fmt.Print(res)
}

func f1(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * f1(n-1)
	}
}

func f2(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return f2(n-1) + f2(n-2)
	}
}
