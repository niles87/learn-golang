package main

import (
	"fmt"
	"runtime"
	"sync"
)

// using goroutines create a program with race condition
func main() {
	counter := 0
	const routines = 100
	var wg sync.WaitGroup

	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			value := counter
			runtime.Gosched()
			value++
			counter = value
			wg.Done()
		}()
		fmt.Println("Go routines", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("End counter", counter)
}
