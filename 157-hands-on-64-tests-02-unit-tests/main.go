package main

import "fmt"

func main() {
	fmt.Println(doOp(2, 3, add))
	fmt.Println(doOp(2, 3, sub))
}

func doOp(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
