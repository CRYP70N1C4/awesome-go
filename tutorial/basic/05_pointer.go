package main

import "fmt"

func main() {
	a := 3
	fmt.Println(a)
	add1(a)
	fmt.Println(a)
	a = add2(&a)// &a 传a的地址
	fmt.Println(a)
}

func add1(a int) int {
	a = a + 1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}
