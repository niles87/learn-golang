package main

import (
	"fmt"
	"runtime"
	"sync"
)

// using previous code fix the race condition using mutex
func main() {
	counter := 0
	const routines = 100
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			mutex.Lock()
			value := counter
			value++
			counter = value
			mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("Go routines", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("End counter", counter)
}
