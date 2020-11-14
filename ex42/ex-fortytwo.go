package main

import (
	"fmt"
)

// create a func that returns a func then call returned func
func main() {
	f := foo()

	f()
}

func foo() func() {
	fmt.Println("Hello from foo")
	return func() {
		fmt.Println("This is returned ;)")
	}
}
