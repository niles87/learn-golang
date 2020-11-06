package main

import "fmt"

type hellogo int

var x hellogo
var y int

// use last exercise code add a new variable of type int
// assign created variable to previous custom typed variable after conversion
// print out the result and type
func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
