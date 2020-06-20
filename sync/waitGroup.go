package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    wg := new(sync.WaitGroup)
    
    for i:= 0; i< 10; i++ {
        wg.Add(1)
        go func(n int) {
            fmt.Println(n)
            wg.Done() // defer wg.Done()
            //fmt.Println(n)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("the end")
}
