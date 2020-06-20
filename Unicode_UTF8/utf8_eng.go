package main

import "fmt"

func main() {
    s := "Hello, world"
    
    fmt.Printf("%c\n", s[0])
    fmt.Printf("%c\n", s[len(s)-1])
}