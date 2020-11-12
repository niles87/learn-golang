package main

import "fmt"

// create a 'person' struct that stores first, last names and three favorite ice cream flavors
type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		iceCream:  []string{"chocolate", "rocky road", "tootie fruity"},
	}
	p2 := person{
		firstName: "Jane",
		lastName:  "Doe",
		iceCream:  []string{"vanilla", "mint chocochip", "daquari ice"},
	}

	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.iceCream {
		fmt.Println(v)
	}

	fmt.Println(p2.firstName, p2.lastName)
	for _, v := range p2.iceCream {
		fmt.Println(v)
	}
}
