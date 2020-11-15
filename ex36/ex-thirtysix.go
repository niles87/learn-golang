package main

import "fmt"

// functions
// create two functions one that takes in a variadic parameter of type int and returns the sum
// one that takes a slice of int and returns the sum then print both
func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f := foo(xi...)
	b := bar(xi)
	fmt.Println("f value:", f, "\nb value:", b)
}

func foo(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func bar(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
