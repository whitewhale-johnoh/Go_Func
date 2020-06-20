package main

import (
    "crypto/sha512"    
    "fmt"
)

func main() {
    s := "Hello, world"
    h1 := sha512.Sum512([]byte(s))
    fmt.Printf("%x\n",h1)
    
    sha := sha512.New()
    sha.Write([]byte("Hello, "))
    sha.Write([]byte("world"))
    h2 := sha.Sum(nil)
    fmt.Printf("%x\n",h2)
}