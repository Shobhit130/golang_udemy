package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

//	type Stringer interface {
//	    String() string
//	}
//
// Stringer interface has a method named String
// we implement this method in book struct so book is now of type Stringer
// basically book now implements the Stringer interface
func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

// this also implements Stringer interface
type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "West with the night",
	}

	var n count = 42

	//And because these have implemented the stringer interface, it's going to change the way they print.
	//And Println takes into consideration if you are also of Type stringer, It's going to change how it prints.
	fmt.Println(b)
	fmt.Println(n)
}
