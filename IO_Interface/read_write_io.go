package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.OpenFile("hello.txt",os.O_CREATE|os.O_RDWR|os.O_TRUNC,os.FileMode(0644))
    
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    
    r := bufio.NewReader(file)
    w := bufio.NewWriter(file)
    
    rw := bufio.NewReadWriter(r, w)
    
    rw.WriteString("Hello, world!")
    rw.Flush()
    
    file.Seek(0, os.SEEK_SET)
    data, _, _ := rw.ReadLine()
    fmt.Println(string(data))
}