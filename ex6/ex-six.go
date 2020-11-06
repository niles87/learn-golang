package main

import "fmt"

// Write a program that prints a number in decimal, binary, and hex
func main() {
	num := 64
	fmt.Printf("%T\t%d\t%b\t%#x\n", num, num, num, num)
}
