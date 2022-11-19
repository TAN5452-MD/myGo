package main

import "fmt"

func main() {
	//切片是引用类型
	var arr = []int{1, 2, 3, 4, 5}
	//浅拷贝
	//s1 := arr
	//
	var s2 = make([]int, 5)
	//深拷贝
	copy(s2, arr)

	getAll(arr)
}

func add(arr []int) {
	arr = append(arr, 200)
	arr = append(arr, 300)
	fmt.Printf("%v", arr)
}
func add2(arr []int) {
	arr = append(arr, arr[:2]...)
	fmt.Printf("%v", arr)
}
func remove(arr []int) {
	arr = append(arr[:2], arr[3:]...)
	fmt.Printf("%v", arr)
}
func getAll(arr []int) {
	for i, i2 := range arr {
		fmt.Printf("%v,%v \n", i, i2)
	}
}
