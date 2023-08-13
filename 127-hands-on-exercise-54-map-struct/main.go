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
	mp := map[string]person{
		p1.last: p1,
		p2.last : p2,
	}

	for _, val := range p1.ice {
		fmt.Println(val)
	}
	for _, val := range p2.ice {
		fmt.Println(val)
	}

	for _,val := range(mp){
		fmt.Println(val)
	}
}
