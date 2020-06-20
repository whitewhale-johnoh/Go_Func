package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	//PushFront, PushBackList, PushFrontList, InsertAfter, InsertBefore, MoveAfter, MoveBefore, MoveToBack,MoveToFront,Len,Remove

	fmt.Println("Front ", l.Front().Value)
	fmt.Println("Back ", l.Back().Value)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
