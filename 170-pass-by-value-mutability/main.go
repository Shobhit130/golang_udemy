package main

import "fmt"

func intDelta(n *int){
	*n = 43
}

func sliceDelta(x1 []int){
	x1[0] = 99
}

func mapDelta(md map[string]int, s string) {
	md[s] = 33
}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	x1 := []int{1,2,3,4,5}
	fmt.Println(x1)
	sliceDelta(x1)
	fmt.Println(x1)

	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])
}