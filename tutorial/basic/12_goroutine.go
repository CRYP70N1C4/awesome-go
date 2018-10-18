package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(content string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(content)
	}
}

func f1() {
	go say("hello")
	say("world")
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range (a) {
		total += v
	}
	fmt.Println("send")
	c <- total // send total to c
}

func f2() {
	a := []int{2, 5, 6, 8, 9}
	c := make(chan int)
	go sum(a, c)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func f3() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func f4() {
	c := make(chan int, 10)
	for i := 0; i < cap(c); i++ {
		c <- i
	}
	close(c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return

		}
	}
}

func f5() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 10; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func f6() {
	f := func(a, b int) {
		fmt.Println(a + b)
	}

	for i := 0; i < 10; i++ {
		go f(i, i)
	}

	time.Sleep(100)
}

func main() {
	f6()
}
