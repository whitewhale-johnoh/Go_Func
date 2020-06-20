package main

import (
    "fmt"
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(1)
    
    done := make(chan bool, 1)
    count := 12
    
    go func() {
        for i := 0; i < count; i++{
            done <- true
            fmt.Println("go routine : ", i)
        }
    }()
    
    for i := 0; i < count; i++ {
        <-done
        fmt.Println("Main Function : ", i)
    }
}