package main

import "fmt"

func main() {
	i := increment()
	j := increment()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(j())
	fmt.Println(j())
}

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
