package main

import (
	"fmt"
)

func main() {

	var s []int
	fmt.Println("emp:", s, len(s), cap(s), s == nil)

	var b = make([]int, 5)
	fmt.Println("len(b):", len(b), "cap(b):", cap(b), b)

	c := make([]int, 0, 5)
	fmt.Println("len(c):", len(c), "cap(c):", cap(c), c)

	d := b[:2]
	fmt.Println("len(d):", len(d), "cap(d):", cap(d), d)

	e := b[2:5]
	fmt.Println("len(e):", len(e), "cap(e):", cap(e), e)
}
