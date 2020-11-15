package main

import (
	"fmt"
	"math"
)

// functions
type circle struct {
	radius float64
}
type square struct {
	length float64
}

type shape interface {
	area() float64
}

func main() {
	c := circle{
		radius: 5.0,
	}

	s := square{
		length: 6.0,
	}
	info(c)
	info(s)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * 2.0
}

func (s square) area() float64 {
	return s.length * s.length
}

func info(s shape) {
	fmt.Println(s.area())
}
