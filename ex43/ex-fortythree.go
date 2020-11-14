package main

import "fmt"

// create a func that uses a callback
func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	o := odd(product, sl...)
	p := product(sl...)
	fmt.Println(p, "\n", o)
}

func product(nums ...int) int {
	p := 1
	for _, v := range nums {
		p *= v
	}
	return p
}

func odd(f func(n ...int) int, nums ...int) int {
	var s []int
	for _, v := range nums {
		if v%2 != 0 {
			s = append(s, v)
		}
	}
	return f(s...)
}
