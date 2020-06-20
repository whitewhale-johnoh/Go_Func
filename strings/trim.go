package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Trim("Hello,~~ world!~~", "~~"))
	fmt.Println(strings.Trim("  Hello, world!   ", " "))
	fmt.Println(strings.TrimLeft("Hello,~~ world!~~", "~~"))
	fmt.Println(strings.TrimRight("~~Hello,~~ world!~~", "~~"))

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}
	var s = "안녕 Hello 고 Go 언어"
	fmt.Println(strings.TrimFunc(s, f))
	fmt.Println(strings.TrimLeftFunc(s, f))
	fmt.Println(strings.TrimRightFunc(s, f))
	fmt.Println(strings.TrimSpace(" Hello, world!   "))
	fmt.Println(strings.TrimSuffix("Hello, world!", "orld!"))
	fmt.Println(strings.TrimSuffix("Hello, world!", "wor"))
}
