package main

import "fmt"

func main() {
	// Declare a pointer variable
	var p *int

	// Declare an int variable
	var i int = 42

	// Initialize the pointer with the memory address of the int variable
	p = &i

	// Read i through the pointer
	fmt.Println(*p) // 42

	// Change the value of i through the pointer
	*p = 21

	// Read i
	fmt.Println(i) // 21
}
