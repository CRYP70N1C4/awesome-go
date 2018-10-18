package main


import (
	"errors"
	"fmt"
)

func main() {
	basic2()
}

func basic1() {
	fmt.Println("hello world")

	var x complex64 = 1 - 1i
	var y complex64 = 1 + 1i
	var z = x * y
	fmt.Println(z)
	const PI = 3.1415
	var isActive bool = false
	fmt.Println(isActive)

	var s = "hello"
	s += " world"
	fmt.Println(s)

	err := errors.New("submit error ")
	fmt.Println(err)
}

const (
	x = iota
	y = iota
	z = iota
	w
)

const v = iota

const (
	h, i, j = iota, iota, iota
)

func basic2() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
	fmt.Println(v)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
}
