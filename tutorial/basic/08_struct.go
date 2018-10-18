package main

import "fmt"

func main() {
	var P person
	P.name, P.age = "mike", 27
	fmt.Println(P)
	var bob = person{age: 30, name: "bob"}
	fmt.Println(bob)

	var jack = student{person{age: 11, name: "jack"}, "hust"}
	fmt.Println(jack)

}

type person struct {
	name string
	age  int
}

type student struct {
	person
	school string
}
