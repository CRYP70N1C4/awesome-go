package main

import "fmt"

func main() {
	flow4(1)
	flow4(3)
}

func flow1(x int) {
	if x > 10 {
		fmt.Println("bigger than 10")
	} else {
		fmt.Println("less than 10")
	}
}

func flow2() {
	i := 0
Here:
	fmt.Println(i)
	i++
	goto Here
}

func flow3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for sum := 1; sum < 1000; {
		sum += sum
		fmt.Println(sum)
	}
}

func flow4(i int) {
	switch i {
	case 1:
		fmt.Println("ok")
		fallthrough
	case 2, 3, 4:
		fmt.Println("not ok")
	default:
		fmt.Println("bad")
	}

}
