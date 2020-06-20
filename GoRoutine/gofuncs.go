package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r)) // random duration for 0~100
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i) // create 100 goroutine
	}

	fmt.Scanln()
}
