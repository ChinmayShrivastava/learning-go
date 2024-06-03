package main

import "fmt"

func main() {
	i := 10

	j := "j for joker"

	if i > 10 {
		fmt.Println("i is greater than 10")
	} else {
		fmt.Println("i is less than or equal to 10")
	}

	if j == "j for joker" {
		fmt.Println("j is joker")
	} else {
		fmt.Println("j is not joker")
	}
}
