package main

import "fmt"

func main(){
	a := []int{0,1,2,3,4,5}
	
	b := make([]int,6)
	copy(b,a)
	fmt.Println("a ",a)
	fmt.Println("b ",b)
	
	a[0] = 7 //b will remain same as the values are getting copied to a new location (b is pointing to a different location)
	
	fmt.Println("a ",a)
	fmt.Println("b ",b)
	
	//length is 50 and capacity is also 50
	//automatically 0 will be allocated for 50 places
	x1 := make([]int,50) 
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))
	//length is 0 and capacity is 50
	//no values will be assigned initially
	x2 := make([]int,0,50)
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))
}