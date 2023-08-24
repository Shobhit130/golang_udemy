package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hi I am", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "Shobhit",
	}

	// saySomething(p) //cannot pass value of type person
	saySomething(&p) //can pass a value of type *person 
}
