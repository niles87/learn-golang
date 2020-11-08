package main

import "fmt"

// print unicode and capital letter
func main() {
	for i := 65; i < 91; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}
