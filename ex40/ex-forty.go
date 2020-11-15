package main

import "fmt"

// functions
// create and use an anonymous function
func main() {
	fmt.Println("hello from main")
	func() {
		fmt.Println("hello from anonymous")
	}()
}
