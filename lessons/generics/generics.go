package main

import "fmt"

// Define a generic function that takes a map of any type and prints all the key-value pairs
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	// Create a map of string to int
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// Call the generic function with the map
	keys := MapKeys(m)

	// Print the keys
	fmt.Println(keys)
}
