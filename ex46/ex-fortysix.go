package main

import "fmt"

// create a person struct
type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Woods",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.first = "John"
	(*p).last = "Snow"
}
