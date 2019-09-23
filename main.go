package main

import (
	"fmt"
)

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

//this is a method that becomes tied into secretAgent
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

// func (r reciever) indentifier(parameters) (return(s)) {code}

func main() {
	sa := secretAgent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}

	sb := secretAgent{
		person: person{
			"ari",
			"jenssen",
		},
		ltk: true,
	}
	fmt.Println(sa)
	defer sa.speak()
	sb.speak()

}