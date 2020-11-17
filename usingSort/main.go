package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 1, 7, 99, 32, 100}
	xs := []string{"John", "Jane", "Max", "Allen", "Ashley"}

	fmt.Println(xi, xs)

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)
}
