package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("i am,", p.first, p.last)
}

//connecting types that have method speak connected to them aka secretAgent
//a value can be of more than one type
type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("this is a person")
	case secretAgent:
		fmt.Println("this is a secret agent")
	}
	fmt.Println("I called human", h)
}

func main() {
	sa := secretAgent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "doctor",
		last:  "yes",
	}

	sa.speak()
	v := sa.person.first
	fmt.Printf("%T\n", v)
	 z := string(sa.person.first)
	 fmt.Println(z)
	 fmt.Printf("%T\n", z)
	p1.speak()
	fmt.Printf("%T\n", sa)
	fmt.Println(p1)
	bar(p1)
	bar(sa)


	//polymorphism allows method to take in multiple differnt types
}
