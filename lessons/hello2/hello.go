package main

import (
	"fmt"

	"example.com/gogreet"
)

func main() {
	// Get a greeting message and print it.
	message := gogreet.Hello("Gladys")
	fmt.Println(message)
}
