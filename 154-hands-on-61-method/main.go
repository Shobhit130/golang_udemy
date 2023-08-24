package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi my name is", p.first, "and I am", p.age, "years old")
}

func main() {
	p := person{
		first: "Shobhit",
		age:   22,
	}

	p.speak()
}
