package main

import "fmt"

func main(){
	a := []int{0,1,2,3,4}
	b := a

	fmt.Println("a ",a)
	fmt.Println("b ",b)
	
	a[0] = 7 //b also gets changed (a and b both are referencing to the same address)
	
	fmt.Println("a ",a)
	fmt.Println("b ",b)
}