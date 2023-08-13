package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 40
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and is greater than or equal to x which is %v", z, x)
	} else {
		fmt.Printf("z is %v and is less than x which is %v", z, x)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	if _, isExists := m["c"]; isExists {
		fmt.Println("\nYes")
	} else {
		fmt.Println("\nNo")
	}

	arr := []int{1, 2, 3, 4, 5, 6}

	for index, val := range arr {
		fmt.Printf("%v \t %v\n", index, val)
	}
	
}
