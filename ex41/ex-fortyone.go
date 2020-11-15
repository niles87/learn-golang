package main

import "fmt"

// functions
// assign func to variable then call
func main() {
	f := func() {
		fmt.Println("this is f in main")
	}

	f()
}
