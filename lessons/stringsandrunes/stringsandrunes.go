package main

import "fmt"

func main() {

	const somestring = "Hello, 世界"

	// Print the string as a byte slice
	fmt.Printf("Byte slice: %v\n", []byte(somestring))

	// Print the string as a rune slice
	runes := []rune(somestring)

	fmt.Printf("Rune slice: %v\n", runes)

	for i, r := range somestring {
		fmt.Printf("Character %d: %c\n", i, r)
	}

	for i, r := range runes {
		fmt.Printf("Rune %d: %c\n", i, r)
	}

}
