package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// type set interface
type myNumbers interface {
	constraints.Integer | constraints.Float
}

// type constraint - what types this generic function can take in
func addT[T myNumbers](a, b T) T {
	return a + b
}

// type alias
type myAlias int

func main() {
	var a myAlias = 22

	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))

	fmt.Println(addT(a, 2))
	fmt.Println(addT(1.2, 2.2))
}
