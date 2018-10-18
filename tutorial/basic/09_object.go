package main

import (
	"fmt"
	"math"
)

func main() {
	r := Rectangle{10, 20}
	c := Circle{10}
	fmt.Println(r.area())
	fmt.Println(c.area())
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
