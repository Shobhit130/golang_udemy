package main

import "fmt"

func main() {
	foo()
	foo1("Shobhit")
	s1 := foo2("Shobhit")
	fmt.Println(s1)

	s2, s3 := foo3("Shobhit", 22)

	fmt.Println(s2, s3)

	s := foo4([]int{1,2,3,4,5})
	fmt.Println(s)
}

// no params,no returns
func foo() {
	fmt.Println("I am from foo")
}

//1 param, no returns
func foo1(s string) {
	fmt.Println("I am ", s)
}

//1 param, 1 return
func foo2(s string) string {
	s1 := "I am " + s
	return s1
}

//2 params, 2 returns
func foo3(name string, age int) (string, int) {
	s1 := "I am " + name
	a := age * 7
	return s1, a
}

//named returns
func foo4(ii []int)(total int){
	total = 0
	for _,v := range(ii){
		total += v
	}
	return
}