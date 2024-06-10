package main

import "fmt"

func main() {

	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"

	// Iterate over the map
	for key, value := range m {
		fmt.Printf("%d -> %s\n", key, value)
	}

	// Check if a key exists
	if value, ok := m[2]; ok {
		fmt.Printf("The value of key 2 is %s\n", value)
	} else {
		fmt.Println("The key 2 does not exist")
	}

	value, ok := m[1]
	fmt.Println("The value of key 1 is", value)
	fmt.Println("The key 1 exists:", ok)

	fmt.Println(m)

}
