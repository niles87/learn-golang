package main

import "fmt"

// using previous map add a record print out the whole map
func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	m["job_odd"] = []string{"Top hat", "Throwing", "Sneaking"}
	for k, v := range m {
		fmt.Println(k, v)

	}
}
