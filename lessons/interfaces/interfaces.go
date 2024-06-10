package main

type Coffee interface {
	default_size() string
	default_milk() string
}

type Espresso struct {
	size string
	milk string
}

type OatMilkLatte struct {
	size string
	milk string
}

func (e *Espresso) default_size() string {
	e.size = "small"
	return e.size
}

func (e *Espresso) default_milk() string {
	e.milk = "none"
	return e.milk
}

func (o *OatMilkLatte) default_size() string {
	o.size = "large"
	return o.size
}

func (o *OatMilkLatte) default_milk() string {
	o.milk = "oat"
	return o.milk
}

func info(c Coffee) string {
	return c.default_size() + " " + c.default_milk()
}

func main() {
	espresso := Espresso{}
	oatMilkLatte := OatMilkLatte{}

	println(info(&espresso))
	println(info(&oatMilkLatte))
}
