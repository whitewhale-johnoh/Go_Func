package main

import (
        "container/ring"
        "fmt"
)

func main() {
    data := []string{"Maria","John","Andrew","James"}
    
    r := ring.New(len(data))        //create ring assinging the number of nodes
    for i := 0; i < r.Len(); i++ {  // countinuous the number of ring nodes
        r.Value = data[i]           // input value in ring node
        r = r.Next()                // move to next node
    }
    
    r.Do(func(x interface{}){       // trace all ring node
        fmt.Println(x)
    })
    
    fmt.Println("Move Forward :")
    r = r.Move(1)
    
    fmt.Println("Curr : ", r.Value)
    fmt.Println("Prev : ", r.Prev().Value)
    fmt.Println("Next : ", r.Next().Value)
    
}