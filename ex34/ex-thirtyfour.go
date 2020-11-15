package main

import "fmt"

// structs
// create an anonymous struct
func main() {
	s := struct {
		name   string
		cars   []string
		places map[string]bool
	}{
		name: "John Doe",
		cars: []string{
			"skyline",
			"viper",
			"brz",
		},
		places: map[string]bool{
			"school": false,
			"work":   false,
			"home":   true,
		},
	}
	fmt.Println(s)
}
