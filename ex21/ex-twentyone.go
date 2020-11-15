package main

import "fmt"

// grouping data
// create a composite literal array with a length of 5 and assign values for each position
// range over array and print values then print the type of array
func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
}
