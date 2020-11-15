package main

import "fmt"

// programming fundamentals
// create typed and untyped constants

const (
	typed   int = 33
	untyped     = "hello"
)

func main() {
	fmt.Printf("%T\t%v\n", typed, typed)
	fmt.Printf("%T\t%v\n", untyped, untyped)
}
