package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "John",
		Last:  "Doe",
		Age:   44,
	}
	p2 := person{
		First: "Jane",
		Last:  "Smith",
		Age:   28,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	byteSlice, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteSlice))

	unMarshaled()
}

func unMarshaled() {
	type human struct {
		First string `json:"First"`
		Last  string `json:"Last"`
		Age   int    `json:"Age"`
	}
	s := `[{"First":"John","Last":"Doe","Age":44},{"First":"Jane","Last":"Smith","Age":28}]`
	byteSlice := []byte(s)
	fmt.Printf("%T\n%T\n", s, byteSlice)

	var h []human
	err := json.Unmarshal(byteSlice, &h)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("humans")
	for _, v := range h {
		fmt.Println(v.First, v.Last, v.Age)
	}
}
