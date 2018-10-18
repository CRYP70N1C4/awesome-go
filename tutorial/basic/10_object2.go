package main

import "fmt"

func main() {
	mark := Student{Human{"mark", 22, "124"}, "MIT"}
	jack := Employee{Human{"jack", 22, "124"}, "google"}
	mark.SayHi()
	jack.SayHi()
}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("name %s,phone %s \n", h.name, h.phone)
}

func (h *Employee) SayHi() {
	fmt.Printf("Employee name %s,company %s \n", h.name, h.company)
}
