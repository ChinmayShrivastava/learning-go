package main

import "fmt"

type WallColor int

const (
	BRIGHT = iota
	DARK
	DULL
)

var wallColor = map[WallColor]string{
	BRIGHT: "white",
	DARK:   "black",
	DULL:   "gray",
}

func (wc WallColor) String() string {
	return wallColor[wc]
}

func main() {
	color := BRIGHT
	println(fmt.Println(color))

	color = DARK
	println(fmt.Println(color))

	color = DULL
	println(fmt.Println(color))

	color = 4
	println(fmt.Println(color))
}
