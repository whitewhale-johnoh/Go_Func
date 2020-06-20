package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"Name" tag2:"이름"`
	age  int    `tag1:"Age" tag2:"나이"`
}

func main() {
	p := Person{}
	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))

}
