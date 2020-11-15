package main

import "fmt"

// functions
// use defer keyword
func main() {
	defer deferred()
	fmt.Println("Hello from main!")
}

func deferred() {
	fmt.Println("\nI am defered")
}
