package main

import "fmt"

// method sets re enforcement
type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		name: "John Wick",
		age:  47,
	}
	fmt.Println(p1)

	saySomething(&p1)
}

func (p *person) speak() {
	fmt.Println("hello how are you")
}

func saySomething(h human) {
	h.speak()
}
