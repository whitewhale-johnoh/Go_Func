package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Author struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

type Comment struct {
    Id uint64 `json:"id"`
    Author Author `json:"author"`
    Content string `json:"content"`
}

type Article struct {
    Id uint64 `json:"id"`
    Title string `json:"title"`
    Author Author `json:"author"`
    Content string `json:"content"`
    Recommends []string `json:"recommends"`
    Comments []Comment `json:"comments"`
}

func main() {
    data := make([]Article, 1)
    
    data[0].Id=1
    data[0].Title="Hello,world!"
    data[0].Author.Name="Maria"
    data[0].Author.Email="maria@example.com"
    data[0].Content = "Hello~"
    data[0].Recommends = []string{"John","Andrew"}
    data[0].Comments = make([]Comment,1)
    data[0].Comments[0].Id = 1
    data[0].Comments[0].Author.Name = "Andrew"
    data[0].Comments[0].Author.Email = "andrew@hello.com"
    data[0].Comments[0].Content = "Hello Maria"
    
    doc, _ := json.MarshalIndent(data,"","  ")
    err := ioutil.WriteFile("./articles.json",doc, os.FileMode(0644))
    
    if err != nil {
        
        fmt.Println(err)
        return
        
    }
}