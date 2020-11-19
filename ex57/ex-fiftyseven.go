package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// using code from ex54 fix race condition with atomic package
func main() {
	var counter int32
	var wg sync.WaitGroup

	const routines = 100
	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			loaded := atomic.LoadInt32(&counter)
			fmt.Println(loaded)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("End counter", counter)
}
