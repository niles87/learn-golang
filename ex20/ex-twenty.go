package main

import "fmt"

// switch statement checking value of a string
func main() {
	favSport := "wnsl"
	switch favSport {
	case "nfl":
		fmt.Println("football")
	case "mlb":
		fmt.Println("baseball")
	case "wnsl":
		fmt.Println("womans soccer")
	}
}
