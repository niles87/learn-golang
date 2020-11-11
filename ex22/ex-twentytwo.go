package main

import "fmt"

// Using composite literal create a slice of type int and assign 10 values
// range over slice and print values then print type of slice
func main() {
	s := []int{5, 3, 2, 6, 1, 6, 7, 10, 49, 100}
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", s)
}
