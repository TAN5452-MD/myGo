package main

import "fmt"

type Person struct {
	name string
}

func main() {
	per := &Person{
		name: "tom",
	}
	per.showPerson()
	fmt.Print(per.name)
}
func (per *Person) showPerson() {
	//地址传递 自动解引用
	per.name = "alice"

}
