package main

import (
    "encoding/json"
    "fmt"
)

type Author struct {
    Name string
    Email string
}

type Comment struct {
    Id uint64
    Author Author
    Content string
}

type Article struct {
    Id uint64
    Title string
    Author Author
    Content string
    Recommends []string
    Comments []Comment
}

func main() {
    doc := `
    [{
        "id": 1,
        "Title": "Hello, world!",
        "Author" : {
            "Name" : "Maria",
            "Email": "maria@example.com"
        },
        "Content":"Hello~",
        "Recommends": [
            "John",
            "Andrew"
        ],
        "Comments": [{
            "id": 1,
            "Author":{
                "Name":"Andrew",
                "Email":"andrew@hello.com"
            },
            "Content" : "Hello Maria"
            
        }]
        
    }]
    `
    
    var data []Article
    
    json.Unmarshal([]byte(doc),&data)
    fmt.Println(data)

}