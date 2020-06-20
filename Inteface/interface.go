package main

import "fmt"

type MyInt int

func (i MyInt) Print() {
	fmt.Println(i)
}

type Printer interface { //인터페이스는 매서드의 집합
	Print()
}

func main() {
	var i MyInt = 5

	var p Printer

	p = i
	p.Print()
}
