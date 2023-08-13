package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	//anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Shobhit",
		last:  "Agrawal",
		age:   22,
	}

	p2 := person{
		first: "Eshaan",
		last: "Moha",
		age: 22,
	}

	fmt.Printf("%T\n",p1)
	fmt.Printf("%T\n",p2)
}
