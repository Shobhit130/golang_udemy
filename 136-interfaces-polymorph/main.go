package main

import "fmt"

//person is of type human also, as it implements speak method
type person struct {
	first string
}

//secretAgent is of type human also, as it implements speak method
type secretAgent struct {
	person
	ltk bool
}
type human interface {
	//if some type has got this method, then it will be of type human
	speak()
}

func saySomething(h human) {
	h.speak()
}

//attach a method to a type person
func (p person) speak() {
	fmt.Println("I am", p.first)
}

//attach a method to a type secretAgent
func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

func main() {
	p1 := person{
		first: "Shobhit",
	}
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	// p1.speak()
	// sa1.speak()

	saySomething(p1)
	saySomething(sa1)
}
