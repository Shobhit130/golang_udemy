package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	//Only the exported (public) fields of a struct will be present in the JSON output. (fields with capital first letter)
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Shobhit",
		Last:  "Agarwal",
		Age:   22,
	}

	p2 := person{
		First: "Eshaan",
		Last:  "Mohapatra",
		Age:   22,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	fmt.Println(string(bs))
}
