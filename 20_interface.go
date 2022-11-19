package main

import "fmt"

type USB interface {
	read()
	write()
}

type computer struct {
	name string
}
type mobile struct {
	name string
}

func main() {
	c := computer{
		name: "lenovo",
	}
	c.read()
	c.write()
}
func (c computer) read() {
	fmt.Printf("%v read...\n", c.name)

}
func (c computer) write() {
	fmt.Printf("%v write...\n", c.name)
}
