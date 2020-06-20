package main

import (
    "compress/gzip"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    file, err := os.Open("hello.txt.gz")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    
    r, err := gzip.NewReader(file)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer r.Close()
    
    b, err := ioutil.ReadAll(r)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println(string(b))
}