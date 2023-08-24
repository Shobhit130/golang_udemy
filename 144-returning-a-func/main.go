package main

import "fmt"

func main() {
	fmt.Println(foo())

	y := bar()
	val := y()
	fmt.Println(val)

	fmt.Printf("%T\n",foo)
	fmt.Printf("%T\n",bar)
	fmt.Printf("%T\n",y)

	//OUTPUT of above 3 lines
	//func() int
	// func() func() int
	// func() int
}

func foo() int{
	return 43
}

func bar() func() int {
	//returning an anonymous function
	return func() int {
		return 42
	}
}
