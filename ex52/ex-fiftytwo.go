package main

import (
	"fmt"
	"sync"
)

// create two additional go routines
var wait sync.WaitGroup

func main() {
	wait.Add(2)

	go foo()
	go bar()

	wait.Wait()

	fmt.Println("main is ending")
}

func foo() {
	fmt.Println("Hello from foo")
	wait.Done()
}

func bar() {
	fmt.Println("Good bye from bar")
	wait.Done()
}
