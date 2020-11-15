package main

import "fmt"

// variables and types
// create own type then declare variable and print out the zero value
// print out the type
// assign variable a value then print out result
type hellogo int

var x hellogo

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
