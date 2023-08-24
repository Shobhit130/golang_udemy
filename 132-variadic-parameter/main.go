package main

import "fmt"

func main() {
	arr := []int{1,2,3,3,4}
	//unfurling a slice
	//we need to pass unlimited number of int values
	//so we unfurl the slice, we can not pass a slice completely
	sum("Agrawal",arr...)
	sum("Shobhit",1, 2, 3, 4, 4)
}

//final incoming parameter should be variadic parameter
func sum(str string,i ...int) {
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	s := 0
	for _,val := range(i){
		s += val
	}
	fmt.Println(str,s)
}
