package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//embedded structs
type agent struct {
	person
	first string
	ltk bool
}

func main() {
	sa1 := agent{
		person: person{
			first: "Shobhit",
			last:  "Agrawal",
			age:   22,
		},
		first: "Eshaan",
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.age, sa1.first, sa1.last, sa1.person,sa1.ltk)
	fmt.Println(sa1.person.first)
}

