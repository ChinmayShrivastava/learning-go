package main

import "fmt"

func main() {

	chantest := make(chan string, 3)

	chantest <- "Hello"
	chantest <- "World"
	chantest <- "!"

	fmt.Println(<-chantest)
	fmt.Println(<-chantest)
	fmt.Println(<-chantest)
}
