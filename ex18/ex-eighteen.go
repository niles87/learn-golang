package main

import "fmt"

// control flow
// using previous code add an else if
func main() {
	x := 7
	if x > 8 {
		fmt.Println("x is greater than 8")
	} else if x == 8 {
		fmt.Println("x is eight")
	} else {
		fmt.Println("x is less than 8")
	}
}
