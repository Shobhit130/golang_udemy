package main

import "fmt"

//cannot use short declaration operator outside a function
var w int = 99

func main() {
	fmt.Println("Hello World")
	//with the short declaration type we do not specify the type
	//the compiler figures out the type by itself (implicit type)
	y := 42
	z := 42.0

	fmt.Printf("%v is of type %T\n",y,y)
	fmt.Printf("%v is of type %T\n",z,z)
	
	var m float32 = 32.345
	
	fmt.Printf("%v is of type %T\n",m,m)
	
	//this won't work
	// z = m
	
	//conversion
	z = float64(m)
	fmt.Printf("%v is of type %T\n",z,z)


}
