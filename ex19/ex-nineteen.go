package main

import "fmt"

// control flow
// switch statement with no switch expression
func main() {
	s := 10
	switch {
	case s == 44:
		fmt.Println(s)
	case s == 10:
		fmt.Println(s)
	}
}
