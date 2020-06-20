package main

import "fmt"

func main() {
	fmt.Printf("%3d\n", 1)
	fmt.Printf("%d\n", 12345)
	fmt.Printf("%03\n", 1)
	fmt.Printf("%03d\n", 12345)

	fmt.Printf("%+d, %d\n", 1, -1)
	fmt.Printf("%9f\n", 1.1234567)
	fmt.Printf("%09f\n", 1.1234567)
	fmt.Printf("%.2f\n", 1.1234567)
	fmt.Printf("%9.2f\n", 1234567.0)
}
