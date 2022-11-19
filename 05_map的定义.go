package main

import "fmt"

func main() {
	//map类型 引用类型
	test1()
}
func test1() {
	//map是无序的
	var m1 = map[string]string{"name": "tom"}
	//m2 := make(map[int]int)

	//返回某个键是否存在
	value, ok := m1["nam"]

	fmt.Printf("%v,%v", value, ok)
}
