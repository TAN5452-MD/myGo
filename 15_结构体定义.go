package main

type Person struct {
	id   int
	name int
	age  int
}

func main() {
	var dog struct {
		id   int
		name string
	}
	dog.id = 1
	dog.name = "wawa"
}
