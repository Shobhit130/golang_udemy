package main

import "fmt"

func main() {
	fmt.Println(fact(5))
}

func fact(a int) int {
	if a == 0 || a == 1 {
		return a
	}
	return a * fact(a-1)
}
