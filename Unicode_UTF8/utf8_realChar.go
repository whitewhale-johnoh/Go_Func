package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    b := []byte("안녕하세요")
    
    r,size := utf8.DecodeRune(b)
    fmt.Printf("%c %d\n", r, size)
    
    r, size = utf8.DecodeRune(b[3:])
    fmt.Printf("%c %d\n", r,size)
    
    r, size = utf8.DecodeLastRune(b)
    fmt.Printf("%c %d\n", r, size)
    
    r, size = utf8.DecodeLastRune(b[:len(b)-3])
    
    fmt.Printf("%c %d\n", r, size)
}