package main

import "fmt"

// using the following slice x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// slice x to create the following
/*
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
*/
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a := x[:5]
	b := x[5:]
	c := x[2:7]
	d := x[1:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
