package main

import (
	"fmt"
	"os"
)

func main() {
	throwsPanic(initUser)
	fmt.Println("here")
}

var user = os.Getenv("USER")

func initUser() {
	if user == "" {
		panic("no value for $USER")
	}
	println(" end ")
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			b = true
		}
	}()
	f()
	return
}
