package main

import "fmt"

// 一个类型可以实现多个接口
type Music interface {
	playMusic()
}
type Video interface {
	playVideo()
}
type Mobile struct {
	model string
}

func (m Mobile) playMusic() {
	fmt.Printf("%v music....\n", m.model)
}
func (m Mobile) playVideo() {
	fmt.Printf("%v video....\n", m.model)
}

func main() {
	m := Mobile{
		model: "iqoo",
	}
	m.playMusic()
	m.playVideo()

}

// 多个类型可以是实现一个接口
type Pet interface {
	eat()
}
type Dog struct {
}
type Cat struct {
}

func (dog Dog) eat() {

}
func (cat Cat) eat() {

}

/*var pet  Pet
pet = Dog{}
pet.eat()
pet = Cat{}
pet.eat()*/