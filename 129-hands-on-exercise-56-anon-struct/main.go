package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Shobiht",
		friends: map[string]int{
			"Eshaan": 32,
		},
		favDrinks: []string{"Cola", "Pepsi"},
	}

	fmt.Println(p1)
}
