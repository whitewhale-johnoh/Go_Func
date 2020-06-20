package main

import "fmt"

func main() {
    var s1 string
    s1 = fmt.Sprint(1, 1.1, "Hello, world!")
    fmt.Print(s1)
    
    var s2 string
    s2 = fmt.Sprintln(1, 1.1, "Hello, World")
    
    fmt.Print(s2)
    var s3 string
    s3 = fmt.Sprintf("%d %f %s\n", 1, 1.1, "Hello world")
    fmt.Print(s3)
    
}