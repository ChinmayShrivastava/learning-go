package main

import "fmt"

func main() {

	chantest := make(chan string)

	go func() {
		chantest <- "Hello"
	}()

	fmt.Println(<-chantest)
}
