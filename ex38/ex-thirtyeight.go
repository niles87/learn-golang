package main

import "fmt"

// create a user defined struct, attach a method then call the method

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "John",
		last:  "Wick",
		age:   47,
	}
	p.speak()
}

func (p person) speak() {
	fmt.Printf("Hi my name is %v %v and I am %v years old\n", p.first, p.last, p.age)
}
