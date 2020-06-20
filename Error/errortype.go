package main

import (
	"fmt"
	"log"
	"time"
)

type HelloOneError struct {
	time  time.Time
	value int
}

func (e HelloOneError) Error() string {
	return fmt.Sprintf("%v: %d is not 1.", e.time, e.value)
}

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}

	return "", HelloOneError{time.Now(), n}
}

func main() {
	s, err := helloOne(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)

	s, err = helloOne(2)
	if err != nil {
		log.Fatal(err)
		//fmt.Println(err.(HelloOneError).value) accessing to structs field by type assertion
	}
}
