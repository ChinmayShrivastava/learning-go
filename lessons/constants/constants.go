package main

import "fmt"

const s string = "s for snake"

func main() {
	fmt.Println(s)

	const i = 60
	fmt.Println(i)

	const j, k int = 70, 80
	fmt.Println(j, k)

	const l, m = 90, "m for monkey"
	fmt.Println(l, m)

	const n = "n for nest"
	fmt.Println(n)

	const o = 100
	fmt.Println(o)
}
