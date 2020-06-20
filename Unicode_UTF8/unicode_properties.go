package main

import (
    "fmt"
    "unicode"
)

func main() {
    fmt.Println(unicode.IsGraphic('1'))
    fmt.Println(unicode.IsGraphic('a'))
    fmt.Println(unicode.IsGraphic('한'))
    
    fmt.Println(unicode.IsGraphic('韓'))
    
    fmt.Println(unicode.IsGraphic('\n'))
    
    fmt.Println(unicode.IsLetter('a'))
    fmt.Println(unicode.IsLetter('1'))
    
    fmt.Println(unicode.IsDigit('1'))
    fmt.Println(unicode.IsControl('\n'))
    fmt.Println(unicode.IsMark('\u17c9'))
    
    fmt.Println(unicode.IsPrint('1'))
    fmt.Println(unicode.IsPunct('.'))
    
    fmt.Println(unicode.IsSpace(' '))
    fmt.Println(unicode.IsSymbol('$'))
    
    fmt.Println(unicode.IsUpper('A'))
    fmt.Println(unicode.IsLower('a'))
}
    
    