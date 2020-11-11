package main

import "fmt"

// create a map with a key of type string and a value of type []string
func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	for i, v := range m["bond_james"] {
		fmt.Println(i, v)
	}
	for i, v := range m["moneypenny_miss"] {
		fmt.Println(i, v)
	}
	for i, v := range m["no_dr"] {
		fmt.Println(i, v)
	}

}
