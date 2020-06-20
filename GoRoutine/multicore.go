package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Number of maximum CPU you use
	fmt.Println(runtime.GOMAXPROCS(0))

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) { //anonymous func for running go routine
			fmt.Println(s, n)
		}(i)
	}
	fmt.Scanln()
}
