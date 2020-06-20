package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := []string{"Hello", "world!"}
	fmt.Println(strings.Join(s1, " "))

	s2 := strings.Split("Hello, world!", " ")
	fmt.Println(s2[1])

	s3 := strings.Fields("Hello, world!")
	fmt.Println(s3[1])

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}

	s4 := strings.FieldsFunc("Hello안녕Hello", f)
	fmt.Println(s4)

	fmt.Println(strings.Repeat("Hello", 2))

	fmt.Println(strings.Replace("Hello,world!", "world", "Go", 1))
	fmt.Println(strings.Replace("Hello Hello", "llo", "Go", 2))
}

//func Join
//func Split
//func Fields
//func FieldsFunc
//func Repeat
//func Replace
//func Trim
//func TrimLeft
//func TrimRight
//func TrimFunc
//func TrimLeftFunc
//func TrimRightFunc
//func TrimSpace
//func TrimSuffix
//func NewReplacer
//func Replace
