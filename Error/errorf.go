package main

import (
	"fmt"
	"log"
)

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}

	return "", fmt.Errorf("%d is not 1", n)
}

func main() {

	defer func() {
		// when runtime error occur, run recover func
		s := recover() // error string returns from log.Panic
		fmt.Println(s) // print out error string
	}()

	s, err := helloOne(1) // proper run
	if err != nil {
		log.Fatal(err)
		//log.Panic(err)
		//log.Print(err)
	}

	s, err = helloOne(2) // cuz error
	if err != nil {
		log.Fatal(err) // print out error stirng and exit program
		//log.Panic(err)
		//log.Print(err)
	}

	// underline doesn't run cuz the program exited
	fmt.Println(s)
	fmt.Println("Hello, world!")
}
