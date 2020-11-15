package main

import "fmt"

// programming fundamentals
// create variable then print its decimal binary and hex
// shift the bits to the left once and assign to new variable then print the same as above
func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := (x << 1)
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)

}
