package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

// use last exercise code and assign varibles values
// save variables in a string and print out the result
func main() {
	s := fmt.Sprintf("%v %v %v\n", x, y, z)

	fmt.Println(s)
}
