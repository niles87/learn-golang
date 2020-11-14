package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(nums...)
	fmt.Println("sum", s)
	sOdd := odd(sum, nums...)
	fmt.Println("sum odds", sOdd)
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}

func odd(f func(nums ...int) int, nums ...int) int {
	var sNum []int
	for _, v := range nums {
		if v%2 != 0 {
			sNum = append(sNum, v)
		}
	}

	return f(sNum...)
}
