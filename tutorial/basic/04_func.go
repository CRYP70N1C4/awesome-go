package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AddAndSub(A, B int) (int, int) {
	return A + B, A - B
}

func myFunc(arg ... int) {
	fmt.Println(arg)
}

type testInt func(int) bool

func isOdd(a int) bool {
	return a%2 == 0
}

func isEven(a int) bool {
	return a%2 == 1
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, val := range slice {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

func demo() {
	f := func(x, y int) int {
		return x + y
	}

	fmt.Println(f(10, 30))
	fmt.Println(f(10, 20))
}

func main() {
	fmt.Println(max(1, 3))
	x, y := AddAndSub(10, 4)
	fmt.Println(x, y)
	myFunc(1, 2, 3)
	myFunc(1, 2)
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(filter(array, isEven))
	fmt.Println(filter(array, isOdd))
	demo()
}
