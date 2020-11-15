package main

import "fmt"

// structs
// use previous code and store values in a map with key of last name
// access each value in map and print them out while ranging over the slice
type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Smith",
		iceCream:  []string{"chocolate", "rocky road", "tootie fruity"},
	}

	p2 := person{
		firstName: "Jane",
		lastName:  "Doe",
		iceCream:  []string{"vanilla", "mint chocochip", "daquari ice"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for _, val := range v.iceCream {
			fmt.Println(val)
		}
	}

}
