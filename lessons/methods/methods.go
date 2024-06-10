package main

type Cylinder struct {
	radius float64
	height float64
	volume float64
}

func (c *Cylinder) Volume() float64 {
	if c.volume == 0 {
		c.volume = 3.14 * c.radius * c.radius * c.height
	}
	return c.volume
}

func main() {
	c := Cylinder{radius: 5, height: 10}
	println(c.Volume())
	println(c.volume)
}
