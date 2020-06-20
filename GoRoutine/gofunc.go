package main

import "fmt"

func hello() {
    fmt.Println("Hello, world!")
}

func main() {
    go hello()
    
    fmt.Scanln() // keep main func not being exit
}