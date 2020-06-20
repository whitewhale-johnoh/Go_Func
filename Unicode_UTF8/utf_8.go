package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    var s string = "한"
    fmt.Println(len(s))
    
    var r rune = '한'
    fmt.Println(utf8.RuneLen(r))
    
    var s1 string = "안녕하세요"
    fmt.Println(utf8.RuneCountInString(s1))
}