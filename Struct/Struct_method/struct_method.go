package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (_ Rectangle) dummy() {
	fmt.Println("dummy")
}

func main() {
	rect1 := Rectangle{30, 30}
	rect1.scaleA(10)
	fmt.Println(rect1)

	rect2 := Rectangle{30, 30}
	rect2.scaleB(10)
	fmt.Println(rect2)
}
