package main

import "fmt"

func main() {
    var num1, num2 int
    n, _ := fmt.Scanf("%d, %d",&num1, &num2)
    fmt.Println("number of input", n)
    fmt.Println(num1, num2)
}