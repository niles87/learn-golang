package main

import "fmt"

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
