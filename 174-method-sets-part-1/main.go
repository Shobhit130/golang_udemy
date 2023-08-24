package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("Hi my name is", d.first, "and I am walking")
}
func (d *dog) run() {
	//this will work without dereferencing also as it is a struct -> d.first
	fmt.Println("Hi my name is", (*d).first, "and I am running")
}

func main() {
	d1 := dog{
		first: "ABC",
	}

	d2 := &dog{
		first: "XYZ",
	}

	d1.run()
	d1.walk()

	d2.run()
	d2.walk()
}
