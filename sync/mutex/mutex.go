package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    var data = []int{}
    var mutex = new(sync.Mutex)
    
    go func(){
        for i := 0; i< 10; i++ {
            mutex.Lock()
            data = append(data, 1)
            fmt.Println("test"+string(i))
            mutex.Unlock()
            
            runtime.Gosched()
        }
    }()
    
    go func() {
        for i := 0; i < 10; i++ {
            mutex.Lock()
            data = append(data,1)
            fmt.Println("cat"+string(i))
            mutex.Unlock()
            runtime.Gosched()
        }
    }()
    
    time.Sleep(2* time.Second)
    
    fmt.Println(len(data))
    fmt.Println(runtime.NumCPU()) // CPU 코어 개수
}