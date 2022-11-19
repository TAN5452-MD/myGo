package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("name is null!")
	}
	return &Person{name: name, age: age}, nil
}

func main() {
	per, err := NewPerson("tom", 18)
	if err == nil {
		fmt.Print(*per)
	}
}
