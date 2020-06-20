package main

import (
    "fmt"
    "reflect"
)

type Data struct {
    a, b int
}

func main() { 
   /* var num int = 1
    fmt.Println(reflect.TypeOf(num))
    
    var s string = "Hello, world!"
    fmt.Println(reflect.TypeOf(s))
    
    var f float32 = 1.3
    fmt.Println(reflect.TypeOf(f))
    
    var data Data = Data{1,2}
    fmt.Println(reflect.TypeOf(data))*/
    
    var f float64 = 1.3
    t := reflect.TypeOf(f)
    v := reflect.ValueOf(f)
    
    fmt.Println(t.Name())
    fmt.Println(t.Size())
    fmt.Println(t.Kind() == reflect.Float64)
    
    fmt.Println(t.Kind() == reflect.Int64)
    fmt. Println(v.Type())
    fmt.Println(v.Kind()==reflect.Float64)
    fmt.Println(v.Kind()==reflect.Int64)
    fmt.Println(v.Float())
}