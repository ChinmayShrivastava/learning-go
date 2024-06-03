package main

import "fmt"

func main() {
	var i int = 10

	switch i {
	case 0:
		fmt.Println("i is 0")
	case 10:
		fmt.Println("i is 10")
	default:
		fmt.Println("i is neither 0 nor 10")
	}
}
