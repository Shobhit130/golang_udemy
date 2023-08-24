package main

import "fmt"

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", c.info)
}

func foo(e error) {
	// fmt.Println("foo ran -", e.Error())
	//this also works
	// fmt.Println("foo ran -", e)
	//this won't work
	// fmt.Println("foo ran -", e.info)
	//this will work - assertion
	fmt.Println("foo ran -", e.(customErr).info)
}

func main() {
	c1 := customErr{
		info: "sample error",
	}

	//c1 is of type error now as customErr has Error() method of the error interface
	//so customErr implements Error interface
	foo(c1)
}
