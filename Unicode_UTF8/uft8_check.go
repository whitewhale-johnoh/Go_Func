package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    var b1 []byte = []byte("안녕하세요")
    fmt.Println(utf8.Valid(b1))
    var b2 [] byte = []byte{0xff, 0xf1, 0xc1}
    fmt.Println(utf8.Valid(b2))
    
    var r1 rune = '한'
    fmt.Println(utf8.ValidRune(r1))
    var r2 rune = 0x11111111
    fmt.Println(utf8.ValidRune(r2))
    
    var s1 string = "한글"
    fmt.Println(utf8.ValidString(s1))
    var s2 string = string([]byte{0xff, 0xf1, 0xc1})
    fmt.Println(utf8.ValidString(s2))
}