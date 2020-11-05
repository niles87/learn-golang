package main

import "fmt"

type hellogo int

var x hellogo

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
