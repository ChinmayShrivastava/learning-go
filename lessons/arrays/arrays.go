package main

import "fmt"

func main() {
	var i [5]int
	fmt.Println("emp:", i)

	var j [10]string
	fmt.Println("emp:", j)

	i[4] = 100
	fmt.Println("set:", i)
	fmt.Println("get:", i[4])

	const b = len(i)
	fmt.Println("len:", b)

	c := [3]int{1, 2, 3}
	fmt.Println("dcl:", c)

	var twod [3][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println("2d:", twod)

	d := [10]int{1000, 5: 1001, 1002, 1003, 1004}
	fmt.Println("d:", d)
}
