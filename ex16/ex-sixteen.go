package main

import "fmt"

// control flow
// print remainder of every number between 10 - 100 when divided by 4
func main() {
	for i := 10; i <= 100; i++ {

		fmt.Println(i % 4)

	}
}
