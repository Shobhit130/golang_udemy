package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//short declaration operator
	a := 42
	fmt.Println(a)

	//multiple initializations
	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// b,c,d,e := 0,1,2,3

	// fmt.Println(b,c,d,e)

	//0 will be assigned to g by default
	//use var for 0 value
	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)
	var h int = 44
	fmt.Println(h)
}
