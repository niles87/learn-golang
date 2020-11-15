package main

import "fmt"

// programming fundamentals
// using iota create 4 constants for the last 4 years then print the values
const (
	pastYears = 2016
	four      = pastYears + iota
	three     = pastYears + iota
	two       = pastYears + iota
	one       = pastYears + iota
)

func main() {
	fmt.Println(four, three, two, one)
}
