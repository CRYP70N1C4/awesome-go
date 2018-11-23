package main

import "fmt"

func reverse(x int) int {
	x = check(x)
	if x >= 0 {
		return check(rever(x, 0))
	}
	return check(-rever(-x, 0))
}

func check(x int) int {
	if x > 1<<31-1 || x < -1<<31 {
		return 0
	}
	return x
}

func rever(x, y int) int {
	if x == 0 {
		return y
	} else if x < 10 {
		return 10*y + x
	} else {
		y = 10*y + x%10
		x = x / 10
		return rever(x, y)
	}
}

func main() {
	fmt.Println(reverse(-2147483651))
	fmt.Println(1 << 3)
	fmt.Println(1<<31 - 1)
}
