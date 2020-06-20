package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Index("Hello, world!", "He"))
	fmt.Println(strings.Index("Hello, world!", "wor"))
	fmt.Println(strings.Index("Hello, world!", "ow"))

	fmt.Println(strings.Index("Hello, world!", "eo"))
	fmt.Println(strings.Index("Hello, world!", "f"))

	var c byte = 'd'
	fmt.Println(strings.IndexByte("Hello, world!", c))
	c = 'f'
	fmt.Println(strings.IndexByte("Hello, world!", c))

	var r rune

	r = '언'
	fmt.Println(strings.IndexRune("고 언어", r))

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}
	fmt.Println(strings.IndexFunc("Go 언어", f))
	fmt.Println(strings.IndexFunc("Go Language", f))

	fmt.Println(strings.LastIndex("Hello Hello Hello, world!", "Hello"))
	fmt.Println(strings.LastIndexAny("Hello, world!", "ol"))
	fmt.Println(strings.LastIndexFunc("Go 언어 안녕", f))
}
