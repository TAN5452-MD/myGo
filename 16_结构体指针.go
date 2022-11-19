package main

import "fmt"

// 没初始化的int都是0 string 都是空 go中 NULL=nil
type Per struct {
	id, name, age int
}

func main() {
	//结构体指针
	p := Per{1, 1, 1}
	var myp *Per
	myp = &p
	fmt.Print(*myp)

	//使用new创建指针
	var p_person = new(Per)
	(*p_person).name = 101
	fmt.Print(*p_person)
}
