package main

import "fmt"

func main() {
	//var for zero value
	var i int
	fmt.Println(i)

	//short declaration operator
	j := 10
	fmt.Println(j)

	//multiple initializations
	a,b,c := 1,2.43,"hi"
	fmt.Println(a,b,c)

	//var when specificity is required
	var f float64 = 32.342343
	fmt.Println(f)

	//blank identifier
	d,e,_ := 42,43,44
	fmt.Println(d,e)
}
