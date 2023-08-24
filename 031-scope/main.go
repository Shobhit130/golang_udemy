package main

// the imported package "fmt"
// is in the "file block" scope
// of this file

import (
	"fmt"
	"mymodule/031-scope/furtherexplored"
)

// x is in "package block" scope
var x = 42

func main() {
	//x can be accessed here
	fmt.Println(x)

	printMe()

	// y does not exist here yet
	// so this won't work
	// fmt.Println(y)

	//y is in "block" scope
	y := 42
	fmt.Println(y)

	p1 := Person{
		"Shobhit",
	}

	p1.sayHello()

	// variable "shadowing"
	// x is in "block" scope
	x := 32
	fmt.Println(x)

	// package block x is still the same
	printMe()

	furtherexplored.Fascinating()

	if z := 53; z > 50 {
		fmt.Println(z)
	}
	// you can't access z here
	// fmt.Println(z)
}

func printMe() {
	//x can be accessed here
	fmt.Println(x)
}

// type person is in "package block" scope
type Person struct {
	first string
}

// the method sayHello
// which is attached to VALUES of TYPE person
// is in "package block" scope
func (p Person) sayHello() {
	fmt.Println("Hi, my name is ", p.first)
}
