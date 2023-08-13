package main

import "fmt"

type person struct {
	first string
	last  string
	ice   []string
}

func main() {
	p1 := person{
		first: "Shobhit",
		last:  "Agrawal",
		ice:   []string{"Chocolate", "Strawberry"},
	}

	p2 := person{
		first: "Eshaan",
		last:  "Mohapatra",
		ice:   []string{"Butterscotch", "Vanilla", "Orange"},
	}

	for _, val := range p1.ice {
		fmt.Println(val)
	}
	for _, val := range p2.ice {
		fmt.Println(val)
	}
}
