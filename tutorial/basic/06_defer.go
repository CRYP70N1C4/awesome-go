package main

import "fmt"

func main() {
	a:=func1(3)
	fmt.Println(a)
}

func func1(a int) int {
	fmt.Println("func begin")
	defer fmt.Println("hello world")
	defer fmt.Println("hello")
	defer fmt.Println("world")
	fmt.Println("func call")
	return a * 100
}
