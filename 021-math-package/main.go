package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main(){
	fmt.Println("Hello world")
	//prints a random number between [0,n)
	fmt.Printf("%d\n",rand.Intn(1000))
	
	fmt.Printf("%.2f\n",math.Sqrt(1000))

	//In Go, a name is exported if it begins with a capital letter.
	// Pi, which is exported from the math package.
	fmt.Println(math.Pi)
}