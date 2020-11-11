package main

import "fmt"

// create a slice of a slice of string and store data in a multi dimentional slice
func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	s := [][]string{x, y}
	for i, v := range s {
		fmt.Println(i)
		for idx, val := range v {
			fmt.Println("\t", idx, val)
		}
	}

}
