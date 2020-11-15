package main

import "fmt"

// functions
// create two functions one that returns an int
// one that returns an int and string then print both
func main() {
	x := foo()
	fmt.Println("called foo", x)

	y, z := bar()
	fmt.Println("called bar", y, z)
}

func foo() int {
	x := 3
	y := 4
	return x + y
}

func bar() (int, string) {
	x := foo()
	y := fmt.Sprint("The value of foo:", x)
	return x, y
}
