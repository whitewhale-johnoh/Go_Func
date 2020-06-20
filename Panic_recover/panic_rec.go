package main

import "fmt"

func f() {
	defer func() {
		s := recover() //recover 함수로 런타임 애러 상황 복구
		fmt.Println(s)
	}()

	a := [...]int{1, 2}

	for i := 0; i < 5; i++ { //배열의 크기를 벗어난 접근
		fmt.Println(a[i])
	}
}

func main() {
	f()

	fmt.Println("Hello, world!") //런타임 애러가 발생했지만 recover 함수로 복구되서 정상 실행
}
