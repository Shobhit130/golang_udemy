package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 3, 4}
	sum := foo(arr...)

	fmt.Println(sum)

	sum = bar(arr)
	fmt.Println(sum)
}

func foo(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func bar(i []int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
