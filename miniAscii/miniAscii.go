package main

import "fmt"

func main() {
	for i := 33; i < 123; i++ {
		if (i > 64 && i < 90) || (i > 96 && i < 123) {

			fmt.Printf("%+q\n", i)
		}
	}
}
