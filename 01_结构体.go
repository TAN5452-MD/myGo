package main

import "fmt"

type Website struct {
	Name string
}

func main() {
	site := Website{Name: "seesaw"}
	fmt.Printf("%v", site.Name)
}
