package main

import "fmt"

// control flow
// print years you've been alive with C while loop in golang
func main() {
	born := 1987
	for born <= 2020 {
		fmt.Println(born)
		born++
	}
}
