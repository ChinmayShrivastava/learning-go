package main

func next_integer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
