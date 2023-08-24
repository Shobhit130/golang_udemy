package main

import "fmt"

type person struct {
	first string
}

//attach a method to a type
func (p person) speak() {
	fmt.Println("I am", p.first)
}

func main() {
	p1 := person{
		first: "Shobhit",
	}

	p1.speak()
}
