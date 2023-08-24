package main

import "fmt"

func main() {
	//anonymous function

	//assigning anonymous functions to a variable - function expressions
	x := func() {
		fmt.Println("Anonymous")
	}

	y := func(s string) {
		fmt.Println("Hi I am ", s)
	}

	x()
	y("Shobhit")
}
