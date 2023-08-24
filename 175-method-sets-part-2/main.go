package main

import "fmt"

type dog struct {
	first string
}

//value receiver
func (d dog) walk() {
	fmt.Println("Hi my name is", d.first, "and I am walking")
}

//pointer receiver
func (d *dog) run() {
	//this will work without dereferencing also as it is a struct - d.first
	fmt.Println("Hi my name is", (*d).first, "and I am running")
}

type youngin interface {
	walk()
	run()
}

func youngrun(y youngin) {
	y.run()
	y.walk()
}

func main() {
	d1 := dog{
		first: "ABC",
	}

	d2 := &dog{
		first: "XYZ",
	}

	// d1.run()
	// d1.walk()

	// d2.run()
	// d2.walk()

	// youngrun(d1) //error - (method run has pointer receiver)
	youngrun(&d1)
	youngrun(d2)
}

//when you have a value which is a pointer, a type pointer, 
//you could use "pointer receivers or value receivers" to implement an interface.

//If you have a value which is just a value, not a pointer, you can only use "value receivers" only.

//run isn't being implemented by dog d1 because it's a value

//d2 is a pointer. So it's implementing the interface because with pointer semantics, 
//you could use receivers of type value or type pointer right there.