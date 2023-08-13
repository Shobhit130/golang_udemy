package main

import "fmt"

func main() {
	s1 := []string{"a", "b", "c"}

	fmt.Println(s1)

	//slice initially of length 0 (initially no elements) and capcity 10
	x1 := make([]int, 0, 10)

	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))
	x1 = append(x1,0,1,2,3,4,5,6,7,8,9)
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))
	
	x1 = append(x1,10,11,12,13)
	fmt.Println(x1)
	fmt.Println(len(x1))
	fmt.Println(cap(x1))



}
