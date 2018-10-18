package main

import "fmt"

// interface 中定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象实现了此接口

type Fly interface {
	fly()
}

type Bird struct {
	name string
}

func (b *Bird) fly() {
	fmt.Println(b.name, " is flying")
}

func main() {
	var ele Fly = new(Bird)
	ele.fly()
}
