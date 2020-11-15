package main

import "fmt"

// control flow
// print years alive using for without condition
func main() {
	born := 1987
	for {
		if born > 2020 {
			break
		}
		fmt.Println(born)
		born++
	}
}
