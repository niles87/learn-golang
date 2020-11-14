package main

import "fmt"

type person struct {
	first string
	last  string
}

type student struct {
	person
	inSchool bool
}

type human interface {
	isStudent()
}

func main() {
	p1 := student{
		person: person{
			"John",
			"Doe",
		},
		inSchool: true,
	}
	p2 := person{
		"Jane",
		"Smith",
	}

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
	p1.isStudent()
	p2.isStudent()
	bar(p1)
	bar(p2)
}

func (p person) greeting() string {
	return fmt.Sprint("Hi I'm ", p.first, " ", p.last, ".")
}

func (p person) isStudent() {
	fmt.Println(p.greeting(), "I am not a student")
}

func (s student) isStudent() {
	if s.inSchool {
		fmt.Println(s.greeting(), "I am a student")
		return
	}
	fmt.Println(s.greeting(), "I am not a student")
}

func bar(h human) {
	fmt.Println("I am a human", h)
}
