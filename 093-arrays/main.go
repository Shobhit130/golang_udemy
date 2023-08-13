package main

import "fmt"

func main(){
	fmt.Println("Hello World")

	//array literal
	a := [3]int{1,2,3}
	fmt.Println(a)

	//let the compiler determine the length
	b := [...]string{"Hello","World"}
	fmt.Println(b)

	var c [2]int
	c[0] = 4
	c[1] = 5

	fmt.Print(c)

	fmt.Printf("\n%T\n%T",a,c)

	//not possible
	// b = a
	// c = a

	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(len(c))

}