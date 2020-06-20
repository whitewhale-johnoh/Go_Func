package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

type Student struct {
	p      Person
	school string
	grade  int
}

func main() {
	var s Student
	s.p.greeting()
}
