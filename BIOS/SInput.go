package main

import "fmt"

func main() {
    var num1 int
    var num2 float32
    var s string
    
    input1 := "1\n1.1\nHello"
    n, _ := fmt.Sscan(input1, &num1, &num2, &s)
    
    fmt.Println("입력 개수:", n)
    fmt.Println(num1, num2, s)
    
    input2 := "1 1.1 Hello"
    n, _ = fmt.Sscanln(input2, &num1, &num2, &s)
    
    fmt.Println("입력 개수:", n)
    fmt.Println(num1,num2,s)
    
    input3 := "1,1.1,Hello"
    n, _ = fmt.Sscanf(input3, "%d, %f, %s", &num1, num2, &s)
    
    fmt.Println("입력 개수:",n)
    fmt.Println(num1,num2,s)
}