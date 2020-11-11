package main

import "fmt"

// using previous map delete a record from the map then print out map
func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	if _, ok := m["bond_james"]; ok {
		delete(m, "bond_james")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
