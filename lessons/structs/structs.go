package main

import "fmt"

type Car struct {
	Make  string
	Model string
	Year  int
}

func main() {

	// Create a new car
	car := Car{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2015,
	}

	// Print the car
	fmt.Printf("Car: %v\n", car)

	// print the type of car
	fmt.Printf("Type of car: %T\n", car)

}

// Car: {Toyota Corolla 2015}
