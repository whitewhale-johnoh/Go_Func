package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, world!", "wo"))
	fmt.Println(strings.Contains("Hello, world!", "w o"))
	fmt.Println(strings.Contains("Hello, world!", "ow"))

	fmt.Println(strings.ContainsAny("Hello, world!", "wo"))
	fmt.Println(strings.ContainsAny("Hello, world!", "w o"))
	fmt.Println(strings.ContainsAny("Hello, world!", "ow"))

	fmt.Println(strings.Count("Hello, Helium", "He"))

	//func Contains
	//func ContainsAny
	//func ContainsRune
	//func Count
	//func HasPrefix
	//func HasSuffix
	//func Index
	//func IndexAny
	//func IndexByte
	//func IndexRune
	//func IndexFunc
	//func LastIndex
	//func LastIndexAny
	//func LastIndexFunc
}
